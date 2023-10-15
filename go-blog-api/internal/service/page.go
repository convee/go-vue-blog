package service

import (
	"github.com/convee/go-blog-api/internal/daos"
	"github.com/convee/go-blog-api/internal/models"
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
		pages    []models.Page
		total    int64
		whereMap = make(map[string]interface{})
	)
	if len(req.Title) > 0 {
		whereMap["title like"] = "%" + req.Title + "%"
	}
	db := s.dao.DB
	build, vars, err := daos.WhereBuild(whereMap)
	if err != nil {
		return nil, err
	}
	if len(vars) > 0 {
		db = db.Where(build, vars)
	}
	_ = db.Limit(req.GetPageSize()).Offset(req.GetOffset()).Order("updated_at DESC").Find(&pages).Limit(-1).Offset(-1).Count(&total)
	return models.PageListRes{
		PageInfo: models.PageInfo{
			Page:      req.GetPage(),
			Total:     total,
			PageSize:  req.GetPageSize(),
			TotalPage: req.GetTotalPage(total),
		},
		PageData: pages,
	}, nil
}

func (s *PageService) Detail(ctx *gin.Context, ident string) (interface{}, error) {
	var page models.Page
	_ = s.dao.DB.Where("ident=?", ident).Find(&page)
	return page, nil

}
func (s *PageService) GetPageById(ctx *gin.Context, id string) (interface{}, error) {
	var page models.Page
	_ = s.dao.DB.Where("id=?", id).Find(&page)
	return page, nil
}

func (s *PageService) Add(ctx *gin.Context, req models.PageAddReq) (interface{}, error) {
	var page models.Page
	s.dao.DB.Where("title=?", req.Title).Find(&page)
	if page.Id > 0 {
		return nil, errors.New("页面名称已存在")
	}
	s.dao.DB.Where("ident=?", req.Ident).Find(&page)
	if page.Id > 0 {
		return nil, errors.New("页面标识已存在")
	}
	page.Title = req.Title
	page.Ident = req.Ident
	page.Content = req.Content
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
	s.dao.DB.Model(page).Where("id != ? and title=?", req.Id, req.Title).Count(&count)
	if count > 0 {
		return nil, errors.New("页面名称已存在")
	}

	s.dao.DB.Model(page).Where("id != ? and ident=?", req.Id, req.Ident).Count(&count)
	if count > 0 {
		return nil, errors.New("页面标识已存在")
	}
	page.Title = req.Title
	page.Ident = req.Ident
	page.Content = req.Content
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
