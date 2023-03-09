package service

import (
	"github.com/convee/go-vue-blog/internal/daos"
	"github.com/convee/go-vue-blog/internal/models"
	"github.com/gin-gonic/gin"
)

// CategoryService 页面服务
type CategoryService struct {
	dao *daos.Dao
}

func NewCategoryService() *CategoryService {
	return &CategoryService{
		dao: daos.NewDao(),
	}
}
func (s *CategoryService) List(ctx *gin.Context, req models.CategoryListReq) (interface{}, error) {
	var (
		categories []models.Category
		total      int64
	)
	_ = s.dao.DB.Limit(req.GetLimit()).Offset(req.GetOffset()).Find(&categories).Limit(-1).Offset(-1).Count(&total)
	return models.CategoryListRes{
		PageInfo: models.PageInfo{
			Page:    req.GetPage(),
			Total:   total,
			PerPage: req.GetLimit(),
		},
		Data: categories,
	}, nil
}

func (s *CategoryService) Detail(ctx *gin.Context, id string) (interface{}, error) {
	var category models.Category
	_ = s.dao.DB.Where("id=?", id).Find(&category)
	return category, nil

}

func (s *CategoryService) GetAll(ctx *gin.Context) (interface{}, error) {
	var categories []models.Category
	_ = s.dao.DB.Find(&categories)
	return categories, nil
}
