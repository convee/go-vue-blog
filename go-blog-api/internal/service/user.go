package service

import (
	"github.com/convee/go-blog-api/internal/daos"
	"github.com/convee/go-blog-api/internal/enum"
	"github.com/convee/go-blog-api/internal/models"
	"github.com/gin-gonic/gin"
)

// UserService 用户服务
type UserService struct {
	dao *daos.Dao
}

func NewUserService() *UserService {
	return &UserService{
		dao: daos.NewDao(),
	}
}

func (u *UserService) Index(ctx *gin.Context) *models.AuthInfo {
	user := ctx.MustGet(enum.UserValueAuth).(*models.AuthInfo)
	return user
}
