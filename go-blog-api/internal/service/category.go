package service

import (
	"errors"
	"github.com/convee/go-blog-api/internal/daos"
	"github.com/convee/go-blog-api/internal/models"
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
		whereMap   = make(map[string]interface{})
	)
	db := s.dao.DB
	if len(req.Name) > 0 {
		whereMap["name like"] = "%" + req.Name + "%"
	}
	build, vars, err := daos.WhereBuild(whereMap)
	if err != nil {
		return nil, err
	}
	if len(vars) > 0 {
		db = db.Where(build, vars)
	}
	_ = db.Limit(req.GetPageSize()).Offset(req.GetOffset()).Order("updated_at DESC").Find(&categories).Limit(-1).Offset(-1).Count(&total)
	return models.CategoryListRes{
		PageInfo: models.PageInfo{
			Page:      req.GetPage(),
			Total:     total,
			PageSize:  req.GetPageSize(),
			TotalPage: req.GetTotalPage(total),
		},
		PageData: categories,
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

func (s *CategoryService) Add(ctx *gin.Context, req models.CategoryAddReq) (interface{}, error) {
	var category models.Category
	s.dao.DB.Where("name=?", req.Name).Find(&category)
	if category.Id > 0 {
		return nil, errors.New("名称已存在")
	}
	category.Name = req.Name
	err := s.dao.DB.Create(&category).Error
	if err != nil {
		return nil, err
	}
	return category, nil

}

func (s *CategoryService) Update(ctx *gin.Context, req models.CategoryUpdateReq) (interface{}, error) {
	var (
		category models.Category
		count    int64
	)
	s.dao.DB.Where("id=?", req.Id).Find(&category)
	if category.Id <= 0 {
		return nil, errors.New("不存在该记录")
	}
	s.dao.DB.Model(category).Where("id != ? and name=?", req.Id, req.Name).Count(&count)

	if count > 0 {
		return nil, errors.New("名称已存在")
	}

	category.Name = req.Name
	err := s.dao.DB.Save(&category).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (s *CategoryService) Delete(ctx *gin.Context, req models.CategoryDelReq) (interface{}, error) {
	var (
		category models.Category
		articles []models.Article
	)
	s.dao.DB.Where("id=?", req.Id).Find(&category)
	if category.Id <= 0 {
		return nil, errors.New("不存在该记录")
	}
	s.dao.DB.Where("category_id=?", req.Id).Find(&articles)
	if len(articles) > 0 {
		return nil, errors.New("该分类已被使用")
	}
	err := s.dao.DB.Delete(&category).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}
