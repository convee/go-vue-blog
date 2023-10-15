package api

import (
	"github.com/convee/go-blog-api/internal/pkg/app"
	"github.com/convee/go-blog-api/internal/pkg/e"
	"github.com/convee/go-blog-api/internal/service"
	"github.com/gin-gonic/gin"
)

type TagHandler struct {
	s *service.TagService
}

func NewTagHandler() *TagHandler {
	return &TagHandler{
		s: service.NewTagService(),
	}
}

func (h *TagHandler) List(ctx *gin.Context) {
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
