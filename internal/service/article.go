package service

import (
	"github.com/convee/go-vue-blog/internal/daos"
	"github.com/convee/go-vue-blog/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
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
		whereMap = make(map[string]interface{})
	)
	db := s.dao.DB

	if len(req.Keyword) > 0 {
		whereMap["content like"] = "%" + req.Keyword + "%"
	}
	if req.ArticleId > 0 {
		whereMap["article_id"] = req.ArticleId
	}
	build, vars, err := daos.WhereBuild(whereMap)
	if err != nil {
		return nil, err
	}
	if len(vars) > 0 {
		db.Where(build, vars)
	}
	_ = db.Limit(req.GetLimit()).Offset(req.GetOffset()).Find(&articles).Limit(-1).Offset(-1).Count(&total)
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

func (s *ArticleService) GetAll(ctx *gin.Context) (interface{}, error) {
	var categories []models.Article
	_ = s.dao.DB.Find(&categories)
	return categories, nil
}

func (s *ArticleService) Add(ctx *gin.Context, req models.ArticleAddReq) (interface{}, error) {
	var article models.Article
	s.dao.DB.Where("title=?", req.Title).Find(&article)
	if article.Id > 0 {
		return nil, errors.New("标题已存在")
	}
	article.Title = req.Title
	err := s.dao.DB.Create(&article).Error
	if err != nil {
		return nil, err
	}
	return article, nil

}

func (s *ArticleService) Update(ctx *gin.Context, req models.ArticleUpdateReq) (interface{}, error) {
	var (
		article models.Article
		count   int64
	)
	s.dao.DB.Where("id=?", req.Id).Find(&article)
	if article.Id <= 0 {
		return nil, errors.New("不存在该记录")
	}
	s.dao.DB.Model(article).Where("id != ? and title=?", req.Id, req.Title).Count(&count)
	if count > 0 {
		return nil, errors.New("标题已存在")
	}

	article.Title = req.Title
	err := s.dao.DB.Save(&article).Error
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (s *ArticleService) Delete(ctx *gin.Context, req models.ArticleDelReq) (interface{}, error) {
	var (
		article models.Article
	)
	s.dao.DB.Where("id=?", req.Id).Find(&article)
	if article.Id <= 0 {
		return nil, errors.New("不存在该记录")
	}
	err := s.dao.DB.Delete(&article).Error
	if err != nil {
		return nil, err
	}
	return article, nil
}
