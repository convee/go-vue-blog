package api

import (
	"github.com/convee/go-vue-blog/internal/pkg/app"
	"github.com/convee/go-vue-blog/internal/pkg/e"
	"github.com/convee/go-vue-blog/internal/service"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	s *service.CategoryService
}

func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{
		s: service.NewCategoryService(),
	}
}

func (h *CategoryHandler) List(ctx *gin.Context) {
	var (
		appG = app.Gin{C: ctx}
	)
	all, err := h.s.GetAll(ctx)
	if err != nil {
		appG.Error(e.ERROR, err.Error(), nil)
		return
	}
	appG.Success(all)
}
