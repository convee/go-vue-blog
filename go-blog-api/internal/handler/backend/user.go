package backend

import (
	"github.com/convee/go-blog-api/internal/pkg/app"
	"github.com/convee/go-blog-api/internal/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	s *service.UserService
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		s: service.NewUserService(),
	}
}

func (u *UserHandler) Index(ctx *gin.Context) {
	var (
		ag = app.Gin{C: ctx}
	)
	ag.Success(u.s.Index(ctx))
}
