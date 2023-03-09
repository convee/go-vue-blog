package backend

import (
	"github.com/convee/go-vue-blog/internal/models"
	"github.com/convee/go-vue-blog/internal/pkg/app"
	"github.com/convee/go-vue-blog/internal/pkg/e"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	var (
		appG = app.Gin{C: ctx}
		req  models.PageInfo
	)
	validateErr := app.BindQuery(ctx, &req)
	if len(validateErr) > 0 {
		appG.Error(e.INVALID_PARAMS, validateErr[0], nil)
		return
	}
	appG.Success(nil)
}
