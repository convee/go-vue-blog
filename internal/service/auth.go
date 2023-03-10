package service

import (
	"errors"
	"github.com/convee/go-vue-blog/internal/daos"
	"github.com/convee/go-vue-blog/internal/models"
	"github.com/convee/go-vue-blog/pkg/jwt"
	"github.com/convee/go-vue-blog/pkg/utils"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

// AuthService 认证服务
type AuthService struct {
	dao *daos.Dao
}

func NewAuthService() *AuthService {
	return &AuthService{
		dao: daos.NewDao(),
	}
}

func (s *AuthService) AuthCheck(ctx *gin.Context, req models.LoginReq) (interface{}, error) {
	var user models.User
	s.dao.DB.Where("username=?", req.Username).Find(&user)
	if user.Id == 0 {
		return nil, errors.New("用户不存在")
	}
	password := utils.GenPassword(req.Password, user.Salt)
	if password != user.Password {
		return nil, errors.New("密码不正确")
	}
	claim := jwtgo.StandardClaims{
		Subject:   user.Username,
		ExpiresAt: time.Now().Unix() + 2592000,
	}
	token, err := jwt.GenerateToken(claim)
	if err != nil {
		return nil, err
	}
	return struct {
		Token string `json:"token"`
	}{
		token,
	}, nil
}
