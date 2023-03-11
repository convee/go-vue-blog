package service

import (
	"github.com/convee/go-vue-blog/internal/daos"
	"github.com/convee/go-vue-blog/internal/models"
	"github.com/gin-gonic/gin"
)

// TagService 标签服务
type TagService struct {
	dao *daos.Dao
}

func NewTagService() *TagService {
	return &TagService{
		dao: daos.NewDao(),
	}
}
func (s *TagService) List(ctx *gin.Context, req models.TagListReq) (interface{}, error) {
	var (
		tags  []models.Tag
		total int64
	)
	_ = s.dao.DB.Limit(req.GetPageSize()).Offset(req.GetOffset()).Find(&tags).Limit(-1).Offset(-1).Count(&total)
	return models.TagListRes{
		PageInfo: models.PageInfo{
			Page:      req.GetPage(),
			Total:     total,
			PageSize:  req.GetPageSize(),
			TotalPage: req.GetTotalPage(total),
		},
		Data: tags,
	}, nil
}

func (s *TagService) Detail(ctx *gin.Context, id string) (interface{}, error) {
	var tag models.Tag
	_ = s.dao.DB.Where("id=?", id).Find(&tag)
	return tag, nil
}

func (s *TagService) GetAll(ctx *gin.Context) (interface{}, error) {
	var tags []models.Tag
	_ = s.dao.DB.Find(&tags)
	return tags, nil
}
