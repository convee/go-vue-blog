package api

import (
	"github.com/convee/go-vue-blog/internal/models"
	"github.com/convee/go-vue-blog/internal/pkg/app"
	"github.com/convee/go-vue-blog/internal/pkg/e"
	"github.com/convee/go-vue-blog/internal/service"
	"github.com/gin-gonic/gin"
)

type ArticleHandler struct {
	s *service.ArticleService
}

func NewArticleHandler() *ArticleHandler {
	return &ArticleHandler{
		s: service.NewArticleService(),
	}
}

func (h *ArticleHandler) List(ctx *gin.Context) {
	var (
		ag  = app.Gin{C: ctx}
		req models.ArticleListReq
	)
	validateErr := app.BindQuery(ctx, &req)
	if len(validateErr) > 0 {
		ag.Error(e.INVALID_PARAMS, validateErr[0], nil)
		return
	}
	list, err := h.s.List(ctx, req)
	if err != nil {
		ag.Error(e.ERROR, err.Error(), nil)
		return
	}
	ag.Success(list)
}

func (h *ArticleHandler) Detail(ctx *gin.Context) {
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
