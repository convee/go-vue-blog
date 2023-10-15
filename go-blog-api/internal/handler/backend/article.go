package backend

import (
	"github.com/convee/go-blog-api/internal/models"
	"github.com/convee/go-blog-api/internal/pkg/app"
	"github.com/convee/go-blog-api/internal/pkg/e"
	"github.com/convee/go-blog-api/internal/service"
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

func (h *ArticleHandler) Stat(ctx *gin.Context) {
	var (
		ag = app.Gin{C: ctx}
	)
	ag.Success(h.s.Stat(ctx))
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
	id := ctx.Query("id")
	detail, err := h.s.Detail(ctx, id)
	if err != nil {
		ag.Error(e.ERROR, err.Error(), nil)
		return
	}
	ag.Success(detail)
}

func (h *ArticleHandler) Add(ctx *gin.Context) {
	var (
		ag  = app.Gin{C: ctx}
		req models.ArticleAddReq
	)
	validateErr := app.BindJson(ctx, &req)
	if len(validateErr) > 0 {
		ag.Error(e.INVALID_PARAMS, validateErr[0], nil)
		return
	}
	detail, err := h.s.Add(ctx, req)
	if err != nil {
		ag.Error(e.ERROR, err.Error(), nil)
		return
	}
	ag.Success(detail)
}

func (h *ArticleHandler) Update(ctx *gin.Context) {
	var (
		ag  = app.Gin{C: ctx}
		req models.ArticleUpdateReq
	)
	validateErr := app.BindJson(ctx, &req)
	if len(validateErr) > 0 {
		ag.Error(e.INVALID_PARAMS, validateErr[0], nil)
		return
	}
	detail, err := h.s.Update(ctx, req)
	if err != nil {
		ag.Error(e.ERROR, err.Error(), nil)
		return
	}
	ag.Success(detail)
}

func (h *ArticleHandler) Delete(ctx *gin.Context) {
	var (
		ag  = app.Gin{C: ctx}
		req models.ArticleDelReq
	)
	validateErr := app.BindJson(ctx, &req)
	if len(validateErr) > 0 {
		ag.Error(e.INVALID_PARAMS, validateErr[0], nil)
		return
	}
	detail, err := h.s.Delete(ctx, req)
	if err != nil {
		ag.Error(e.ERROR, err.Error(), nil)
		return
	}
	ag.Success(detail)
}
