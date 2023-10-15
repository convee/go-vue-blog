package service

import (
	"errors"
	"time"

	"github.com/convee/go-blog-api/internal/daos"
	"github.com/convee/go-blog-api/internal/models"
	"github.com/convee/go-blog-api/pkg/jwt"
	"github.com/convee/go-blog-api/pkg/utils"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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
	s.dao.DB.Where("name=?", req.Name).Find(&user)
	if user.Id == 0 {
		return nil, errors.New("用户不存在")
	}
	password := utils.GenPassword(req.Password, user.Salt)
	if password != user.Password {
		return nil, errors.New("密码不正确")
	}
	claim := jwtgo.StandardClaims{
		Subject:   user.Name,
		ExpiresAt: time.Now().Unix() + 2592000,
	}
	token, err := jwt.GenerateToken(claim)
	if err != nil {
		return nil, err
	}
	return models.TokenInfo{
		Token:  token,
		Name:   user.Name,
		Id:     user.Id,
		Avatar: user.Avatar,
	}, nil
}
