package api

import (
	"github.com/convee/go-blog-api/internal/pkg/app"
	"github.com/convee/go-blog-api/internal/pkg/e"
	"github.com/convee/go-blog-api/internal/service"
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
		ag = app.Gin{C: ctx}
	)
	all, err := h.s.GetAll(ctx)
	if err != nil {
		ag.Error(e.ERROR, err.Error(), nil)
		return
	}
	ag.Success(all)
}
