package backend

import (
	"github.com/convee/go-blog-api/internal/models"
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

func (h *CategoryHandler) All(ctx *gin.Context) {
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

func (h *CategoryHandler) List(ctx *gin.Context) {
	var (
		ag  = app.Gin{C: ctx}
		req models.CategoryListReq
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

func (h *CategoryHandler) Detail(ctx *gin.Context) {
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

func (h *CategoryHandler) Add(ctx *gin.Context) {
	var (
		ag  = app.Gin{C: ctx}
		req models.CategoryAddReq
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

func (h *CategoryHandler) Update(ctx *gin.Context) {
	var (
		ag  = app.Gin{C: ctx}
		req models.CategoryUpdateReq
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

func (h *CategoryHandler) Delete(ctx *gin.Context) {
	var (
		ag  = app.Gin{C: ctx}
		req models.CategoryDelReq
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
