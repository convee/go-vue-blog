package api

import (
	"github.com/convee/go-blog-api/internal/pkg/app"
	"github.com/convee/go-blog-api/internal/pkg/e"
	"github.com/convee/go-blog-api/internal/service"
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
	id := ctx.Param("ident")
	detail, err := h.s.Detail(ctx, id)
	if err != nil {
		ag.Error(e.ERROR, err.Error(), nil)
		return
	}
	ag.Success(detail)
}
