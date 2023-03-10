package backend

import (
	"github.com/convee/go-vue-blog/internal/models"
	"github.com/convee/go-vue-blog/internal/pkg/app"
	"github.com/convee/go-vue-blog/internal/pkg/e"
	"github.com/convee/go-vue-blog/internal/service"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	s *service.AuthService
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		s: service.NewAuthService(),
	}
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	var (
		ag  = app.Gin{C: ctx}
		req models.LoginReq
	)
	validateErr := app.BindJson(ctx, &req)
	if len(validateErr) > 0 {
		ag.Error(e.INVALID_PARAMS, validateErr[0], nil)
		return
	}

	token, err := h.s.AuthCheck(ctx, req)
	if err != nil {
		ag.Error(e.ERROR, err.Error(), nil)
		return
	}
	ag.Success(token)
}
