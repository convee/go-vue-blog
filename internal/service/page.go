package service

import (
	"github.com/convee/go-vue-blog/internal/daos"
	"github.com/convee/go-vue-blog/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// PageService 页面服务
type PageService struct {
	dao *daos.Dao
}

func NewPageService() *PageService {
	return &PageService{
		dao: daos.NewDao(),
	}
}
func (s *PageService) List(ctx *gin.Context, req models.PageListReq) (interface{}, error) {
	var (
		pages []models.Page
		total int64
	)
	_ = s.dao.DB.Limit(req.GetLimit()).Offset(req.GetOffset()).Find(&pages).Limit(-1).Offset(-1).Count(&total)
	return models.PageListRes{
		PageInfo: models.PageInfo{
			Page:    req.GetPage(),
			Total:   total,
			PerPage: req.GetLimit(),
		},
		Data: pages,
	}, nil
}

func (s *PageService) Detail(ctx *gin.Context, id string) (interface{}, error) {
	var page models.Page
	_ = s.dao.DB.Where("id=?", id).Find(&page)
	return page, nil

}

func (s *PageService) GetAll(ctx *gin.Context) (interface{}, error) {
	var categories []models.Page
	_ = s.dao.DB.Find(&categories)
	return categories, nil
}

func (s *PageService) Add(ctx *gin.Context, req models.PageAddReq) (interface{}, error) {
	var page models.Page
	s.dao.DB.Where("title=?", req.Title).Find(&page)
	if page.Id > 0 {
		return nil, errors.New("名称已存在")
	}
	page.Title = req.Title
	err := s.dao.DB.Create(&page).Error
	if err != nil {
		return nil, err
	}
	return page, nil

}

func (s *PageService) Update(ctx *gin.Context, req models.PageUpdateReq) (interface{}, error) {
	var (
		page  models.Page
		count int64
	)
	s.dao.DB.Where("id=?", req.Id).Find(&page)
	if page.Id <= 0 {
		return nil, errors.New("不存在该记录")
	}
	s.dao.DB.Where("id != ? and title=?", req.Id, req.Title).Count(&count)
	if count > 0 {
		return nil, errors.New("名称已存在")
	}

	page.Title = req.Title
	err := s.dao.DB.Save(&page).Error
	if err != nil {
		return nil, err
	}
	return page, nil
}

func (s *PageService) Delete(ctx *gin.Context, req models.PageDelReq) (interface{}, error) {
	var (
		page models.Page
	)
	s.dao.DB.Where("id=?", req.Id).Find(&page)
	if page.Id <= 0 {
		return nil, errors.New("不存在该记录")
	}
	err := s.dao.DB.Delete(&page).Error
	if err != nil {
		return nil, err
	}
	return page, nil
}
