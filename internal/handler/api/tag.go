package api

import (
	"github.com/convee/go-vue-blog/internal/pkg/app"
	"github.com/convee/go-vue-blog/internal/pkg/e"
	"github.com/convee/go-vue-blog/internal/service"
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
		appG = app.Gin{C: ctx}
	)
	all, err := h.s.GetAll(ctx)
	if err != nil {
		appG.Error(e.ERROR, err.Error(), nil)
		return
	}
	appG.Success(all)
}
