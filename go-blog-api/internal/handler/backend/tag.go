package backend

import (
	"github.com/convee/go-blog-api/internal/models"
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

func (h *TagHandler) All(ctx *gin.Context) {
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

func (h *TagHandler) List(ctx *gin.Context) {
	var (
		ag  = app.Gin{C: ctx}
		req models.TagListReq
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

func (h *TagHandler) Detail(ctx *gin.Context) {
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

func (h *TagHandler) Add(ctx *gin.Context) {
	var (
		ag  = app.Gin{C: ctx}
		req models.TagAddReq
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

func (h *TagHandler) Update(ctx *gin.Context) {
	var (
		ag  = app.Gin{C: ctx}
		req models.TagUpdateReq
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

func (h *TagHandler) Delete(ctx *gin.Context) {
	var (
		ag  = app.Gin{C: ctx}
		req models.TagDelReq
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
