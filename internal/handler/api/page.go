package api

import (
	"github.com/convee/go-vue-blog/internal/pkg/app"
	"github.com/convee/go-vue-blog/internal/pkg/e"
	"github.com/convee/go-vue-blog/internal/service"
	"github.com/gin-gonic/gin"
)

type PageHandler struct {
	s *service.PageService
}

func NewPageHandler() *PageHandler {
	return &PageHandler{
		s: service.NewPageService(),
	}
}

func (h *PageHandler) Detail(ctx *gin.Context) {
	var (
		ag = app.Gin{C: ctx}
	)
	id := ctx.Param("id")
	detail, err := h.s.Detail(ctx, id)
	if err != nil {
		ag.Error(e.ERROR, err.Error(), nil)
		return
	}
	ag.Success(detail)
}
