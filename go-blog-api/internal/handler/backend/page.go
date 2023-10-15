package backend

import (
	"github.com/convee/go-blog-api/internal/models"
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

func (h *PageHandler) List(ctx *gin.Context) {
	var (
		ag  = app.Gin{C: ctx}
		req models.PageListReq
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

func (h *PageHandler) Detail(ctx *gin.Context) {
	var (
		ag = app.Gin{C: ctx}
	)
	id := ctx.Query("id")
	detail, err := h.s.GetPageById(ctx, id)
	if err != nil {
		ag.Error(e.ERROR, err.Error(), nil)
		return
	}
	ag.Success(detail)
}

func (h *PageHandler) Add(ctx *gin.Context) {
	var (
		ag  = app.Gin{C: ctx}
		req models.PageAddReq
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

func (h *PageHandler) Update(ctx *gin.Context) {
	var (
		ag  = app.Gin{C: ctx}
		req models.PageUpdateReq
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

func (h *PageHandler) Delete(ctx *gin.Context) {
	var (
		ag  = app.Gin{C: ctx}
		req models.PageDelReq
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
