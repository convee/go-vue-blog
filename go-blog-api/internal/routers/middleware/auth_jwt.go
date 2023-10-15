package middleware

import (
	"github.com/convee/go-blog-api/internal/daos"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"net/http"
	"strings"
	"time"

	"github.com/convee/go-blog-api/internal/enum"
	"github.com/convee/go-blog-api/internal/models"
	"github.com/convee/go-blog-api/internal/pkg/app"
	"github.com/convee/go-blog-api/internal/pkg/e"
	"github.com/convee/go-blog-api/pkg/jwt"
	"github.com/convee/go-blog-api/pkg/logger"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type JWTHeader struct {
	Authorization string `header:"Authorization" validate:"required"`
}

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		var ag = app.Gin{C: c}

		code = e.SUCCESS
		var headers JWTHeader
		validateErr := app.BindHeader(c, &headers)
		if len(validateErr) > 0 {
			ag.Error(e.INVALID_PARAMS, "", validateErr)
			c.Abort()
			return
		}
		token := strings.TrimPrefix(headers.Authorization, "Bearer ")
		if token == "" {
			code = e.INVALID_PARAMS
			c.Next()
			return
		}
		claims, err := jwt.ParseToken(token)
		if err != nil {
			code = e.TOKEN_INVALID
			logger.Error("token err", zap.String("token", token), zap.Error(err))
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = e.TOKEN_INVALID
			logger.Error("token expire", zap.String("token", token))
		}
		if code != e.SUCCESS {
			ag.Response(http.StatusUnauthorized, code, data)
			c.Abort()
			return
		}

		var user models.User
		err = daos.GetDB().Where("name", claims.Subject).Find(&user).Error
		if err != nil {
			ag.Response(http.StatusUnauthorized, e.ERROR, nil)
			c.Abort()
			return
		}
		authInfo := &models.AuthInfo{
			Name:   claims.Subject,
			Avatar: user.Avatar,
			Role:   []string{"admin"},
		}
		c.Set(enum.UserValueAuth, authInfo)
		c.Next()
	}

}

func RefreshToken(c *gin.Context) (*models.TokenInfo, error) {
	var headers JWTHeader
	validateErr := app.BindHeader(c, &headers)
	if len(validateErr) > 0 {
		return nil, errors.New("参数错误")
	}
	token := strings.TrimPrefix(headers.Authorization, "Bearer ")
	if token == "" {
		return nil, errors.New("参数错误")

	}
	claim, err := jwt.ParseToken(token)
	if err != nil {
		return nil, errors.New("参数错误")

	}
	newClaim := jwtgo.StandardClaims{
		Subject:   claim.Subject,
		ExpiresAt: time.Now().Unix() + 2592000,
	}
	newToken, err := jwt.GenerateToken(newClaim)
	if err != nil {
		return nil, err
	}
	return &models.TokenInfo{
		Token: newToken,
	}, nil
}
