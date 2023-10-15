package service

import (
	"github.com/convee/go-blog-api/internal/daos"
	"github.com/convee/go-blog-api/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
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
		tags     []models.Tag
		total    int64
		whereMap = make(map[string]interface{})
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
	_ = db.Limit(req.GetPageSize()).Offset(req.GetOffset()).Order("updated_at DESC").Find(&tags).Limit(-1).Offset(-1).Count(&total)
	return models.TagListRes{
		PageInfo: models.PageInfo{
			Page:      req.GetPage(),
			Total:     total,
			PageSize:  req.GetPageSize(),
			TotalPage: req.GetTotalPage(total),
		},
		PageData: tags,
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

func (s *TagService) Add(ctx *gin.Context, req models.TagAddReq) (interface{}, error) {
	var tag models.Tag
	s.dao.DB.Where("name=?", req.Name).Find(&tag)
	if tag.Id > 0 {
		return nil, errors.New("名称已存在")
	}
	tag.Name = req.Name
	err := s.dao.DB.Create(&tag).Error
	if err != nil {
		return nil, err
	}
	return tag, nil

}

func (s *TagService) Update(ctx *gin.Context, req models.TagUpdateReq) (interface{}, error) {
	var (
		tag   models.Tag
		count int64
	)
	s.dao.DB.Where("id=?", req.Id).Find(&tag)
	if tag.Id <= 0 {
		return nil, errors.New("不存在该记录")
	}
	s.dao.DB.Model(tag).Where("id != ? and name=?", req.Id, req.Name).Count(&count)

	if count > 0 {
		return nil, errors.New("名称已存在")
	}

	tag.Name = req.Name
	err := s.dao.DB.Save(&tag).Error
	if err != nil {
		return nil, err
	}
	return tag, nil
}

func (s *TagService) Delete(ctx *gin.Context, req models.TagDelReq) (interface{}, error) {
	var (
		tag         models.Tag
		articleTags []models.ArticleTag
	)
	s.dao.DB.Where("id=?", req.Id).Find(&tag)
	if tag.Id <= 0 {
		return nil, errors.New("不存在该记录")
	}

	s.dao.DB.Where("tag_id=?", req.Id).Find(&articleTags)

	if len(articleTags) > 0 {
		return nil, errors.New("该标签已被使用")
	}
	err := s.dao.DB.Delete(&tag).Error
	if err != nil {
		return nil, err
	}
	return tag, nil
}
