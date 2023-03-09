package service

import (
	"github.com/convee/go-vue-blog/internal/daos"
	"github.com/convee/go-vue-blog/internal/models"
	"github.com/gin-gonic/gin"
)

// ArticleService 文章服务
type ArticleService struct {
	dao *daos.Dao
}

func NewArticleService() *ArticleService {
	return &ArticleService{
		dao: daos.NewDao(),
	}
}
func (s *ArticleService) List(ctx *gin.Context, req models.ArticleListReq) (interface{}, error) {
	var (
		articles []models.Article
		total    int64
	)
	_ = s.dao.DB.Limit(req.GetLimit()).Offset(req.GetOffset()).Find(&articles).Limit(-1).Offset(-1).Count(&total)
	return models.ArticleListRes{
		PageInfo: models.PageInfo{
			Page:    req.GetPage(),
			Total:   total,
			PerPage: req.GetLimit(),
		},
		Data: articles,
	}, nil
}

func (s *ArticleService) Detail(ctx *gin.Context, id string) (interface{}, error) {
	var article models.Article
	_ = s.dao.DB.Where("id=?", id).Find(&article)
	return article, nil

}
