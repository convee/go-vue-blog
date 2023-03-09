package middleware

import (
	"github.com/convee/go-vue-blog/internal/enum"
	"github.com/convee/go-vue-blog/internal/models"
	"github.com/convee/go-vue-blog/internal/pkg/app"
	"github.com/convee/go-vue-blog/internal/pkg/e"
	"github.com/convee/go-vue-blog/pkg/jwt"
	"github.com/convee/go-vue-blog/pkg/logger"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type JWTHeader struct {
	Authorization string `header:"Authorization" validate:"required"`
}

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		var appG = app.Gin{C: c}

		code = e.SUCCESS
		var headers JWTHeader
		validateErr := app.BindHeader(c, &headers)
		if len(validateErr) > 0 {
			appG.Error(e.INVALID_PARAMS, "", validateErr)
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
			appG.Response(http.StatusUnauthorized, code, data)
			c.Abort()
			return
		}
		authInfo := &models.AuthInfo{
			UserId: cast.ToUint64(claims.Subject),
		}
		c.Set(enum.UserValueAuth, authInfo)
		c.Next()
	}
}
