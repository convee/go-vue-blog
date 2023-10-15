package api

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

// @BasePath /api

// List
// @Summary 文章列表
// @Accept json
// @Produce json
// @Param pageNum path int true "页数"
// @Param pageSize path int true "每页数量"
// @Param title query string true "文章标题"
// @Param category_id query int false "分类ID"
// @Success 200 {object} models.ArticleListRes
// @Router /article/list [get]
func (h *ArticleHandler) List(ctx *gin.Context) {
	var (
		ag  = app.Gin{C: ctx}
		req models.FrontArticleListReq
	)
	validateErr := app.BindQuery(ctx, &req)
	if len(validateErr) > 0 {
		ag.Error(e.INVALID_PARAMS, validateErr[0], nil)
		return
	}
	list, err := h.s.GetFrontList(ctx, req)
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
	detail, err := h.s.GetFrontDetail(ctx, id)
	if err != nil {
		ag.Error(e.ERROR, err.Error(), nil)
		return
	}
	ag.Success(detail)
}
