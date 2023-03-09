package service

import (
	"github.com/convee/go-vue-blog/internal/daos"
	"github.com/convee/go-vue-blog/internal/models"
	"github.com/gin-gonic/gin"
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
