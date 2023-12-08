package service

import (
	"github.com/convee/go-blog-api/internal/daos"
	"github.com/convee/go-blog-api/internal/enum"
	"github.com/convee/go-blog-api/internal/models"
	"github.com/convee/go-blog-api/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"gorm.io/gorm"
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

func (s *ArticleService) Stat(ctx *gin.Context) models.StatRes {
	var (
		articleCount  int64
		tagCount      int64
		categoryCount int64
		pageCount     int64
	)
	s.dao.DB.Model(models.Article{}).Count(&articleCount)
	s.dao.DB.Model(models.Tag{}).Count(&tagCount)
	s.dao.DB.Model(models.Category{}).Count(&categoryCount)
	s.dao.DB.Model(models.Page{}).Count(&pageCount)
	return models.StatRes{
		ArticleCount:  articleCount,
		CategoryCount: categoryCount,
		PageCount:     pageCount,
		TagCount:      tagCount,
	}

}
func (s *ArticleService) List(ctx *gin.Context, req models.ArticleListReq) (interface{}, error) {
	user := ctx.MustGet(enum.UserValueAuth).(*models.AuthInfo)
	var (
		articles []models.Article
		total    int64
		whereMap = make(map[string]interface{})
	)
	db := s.dao.DB

	if len(req.Title) > 0 {
		whereMap["title like"] = "%" + req.Title + "%"
	}
	if req.CategoryId > 0 {
		whereMap["category_id"] = req.CategoryId
	}
	build, vars, err := daos.WhereBuild(whereMap)
	if err != nil {
		return nil, err
	}
	if len(vars) > 0 {
		db = db.Where(build, vars)
	}
	_ = db.Limit(req.GetPageSize()).Offset(req.GetOffset()).Order("updated_at DESC").Find(&articles).Limit(-1).Offset(-1).Count(&total)
	var (
		articleList []models.ArticleInfo
		categories  []models.Category
		articleTags []models.ArticleTag
		tags        []models.Tag
		categoryMap = make(map[int]string)
		tagsMap     = make(map[int]string)
	)
	s.dao.DB.Find(&categories)
	s.dao.DB.Find(&tags)

	for _, cate := range categories {
		categoryMap[cast.ToInt(cate.Id)] = cate.Name
	}
	for _, tag := range tags {
		tagsMap[cast.ToInt(tag.Id)] = tag.Name
	}

	for _, article := range articles {
		var category string
		var tagIds []uint
		var tagNames []string
		if _, ok := categoryMap[article.CategoryId]; ok {
			category = categoryMap[article.CategoryId]
		}
		err = s.dao.DB.Where("article_id=?", article.Id).Find(&articleTags).Error
		if err != nil {
			logger.Error("not found article with tag")
		}
		for _, tag := range articleTags {
			tagIds = append(tagIds, tag.TagId)
			tagNames = append(tagNames, tagsMap[cast.ToInt(tag.TagId)])
		}
		articleList = append(articleList, models.ArticleInfo{
			Id:          article.Id,
			Title:       article.Title,
			Content:     article.Content,
			Category:    category,
			Description: article.Desc,
			Author:      user.Name,
			TagIds:      tagIds,
			TagNames:    tagNames,
			CreateDate:  cast.ToString(article.CreatedAt),
			UpdateDate:  cast.ToString(article.UpdatedAt),
		})

	}
	return models.ArticleListRes{
		PageInfo: models.PageInfo{
			Page:      req.GetPage(),
			Total:     total,
			PageSize:  req.GetPageSize(),
			TotalPage: req.GetTotalPage(total),
		},
		PageData: articleList,
	}, nil
}

func (s *ArticleService) Detail(ctx *gin.Context, id string) (interface{}, error) {
	user := ctx.MustGet(enum.UserValueAuth).(*models.AuthInfo)
	var (
		article     models.Article
		category    models.Category
		articleTags []models.ArticleTag
		tags        []models.Tag
		tagIds      []uint
		tagNames    []string
		tagMap      = make(map[uint]string)
	)
	err := s.dao.DB.Where("id=?", id).Find(&article).Error
	if err != nil {
		return nil, err
	}
	if article.Id == 0 {
		return nil, errors.New("文章不存在")
	}
	err = s.dao.DB.Where("article_id=?", id).Find(&articleTags).Error
	if err != nil {
		return nil, err
	}
	err = s.dao.DB.Where("id=?", article.CategoryId).Find(&category).Error
	if err != nil {
		return nil, err
	}
	err = s.dao.DB.Find(&tags).Error
	if err != nil {
		return nil, err
	}
	for _, tag := range tags {
		tagMap[tag.Id] = tag.Name
	}

	for _, articleTag := range articleTags {
		tagIds = append(tagIds, articleTag.TagId)
		if _, ok := tagMap[articleTag.TagId]; ok {
			tagNames = append(tagNames, tagMap[articleTag.TagId])
		}
	}

	return models.ArticleInfo{
		Id:          article.Id,
		Title:       article.Title,
		Content:     article.Content,
		Category:    category.Name,
		Description: article.Desc,
		Author:      user.Name,
		CreateDate:  cast.ToString(article.CreatedAt),
		UpdateDate:  cast.ToString(article.UpdatedAt),
		TagIds:      tagIds,
		TagNames:    tagNames,
		CategoryId:  article.CategoryId,
	}, nil

}

func (s *ArticleService) Add(ctx *gin.Context, req models.ArticleAddReq) (interface{}, error) {
	var (
		article models.Article
	)
	s.dao.DB.Where("title=?", req.Title).Find(&article)
	if article.Id > 0 {
		return nil, errors.New("标题已存在")
	}
	tx := s.dao.DB.Begin()
	article.Title = req.Title
	article.Content = req.Content
	article.CategoryId = req.CategoryId
	err := tx.Create(&article).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	for _, tagId := range req.TagIds {
		var articleTag models.ArticleTag
		articleTag.TagId = tagId
		articleTag.ArticleId = cast.ToInt(article.Id)
		err = tx.Create(&articleTag).Error
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}
	tx.Commit()

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

	// 文章、标签关联表事务
	err := s.dao.DB.Transaction(func(tx *gorm.DB) error {
		// 保存文章
		article.Title = req.Title
		article.Content = req.Content
		article.CategoryId = req.CategoryId
		err := tx.Save(&article).Error
		if err != nil {
			return err
		}
		// 保存标签关联表，先删除后增加
		var articleTags []models.ArticleTag
		err = tx.Where("article_id=?", article.Id).Delete(&articleTags).Error
		if err != nil {
			return err
		}
		for _, tagId := range req.TagIds {
			var articleTag models.ArticleTag
			articleTag.TagId = tagId
			articleTag.ArticleId = cast.ToInt(article.Id)
			err = tx.Create(&articleTag).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
	return article, err
}

func (s *ArticleService) Delete(ctx *gin.Context, req models.ArticleDelReq) (interface{}, error) {
	var (
		article     models.Article
		articleTags []models.ArticleTag
	)
	s.dao.DB.Where("id=?", req.Id).Find(&article)
	if article.Id <= 0 {
		return nil, errors.New("不存在该记录")
	}
	// 删除文章和标签关联记录
	err := s.dao.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Delete(&article).Error
		if err != nil {
			return err
		}
		err = tx.Where("article_id=?", article.Id).Delete(&articleTags).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (s *ArticleService) GetFrontList(ctx *gin.Context, req models.FrontArticleListReq) (interface{}, error) {
	var (
		articles    []models.Article
		articleList []models.ArticleInfo
		articleTags []models.ArticleTag
		tags        []models.Tag
		total       int64
		whereMap    = make(map[string]interface{})
		tagsMap     = make(map[int]string)
	)
	db := s.dao.DB

	if len(req.Title) > 0 {
		whereMap["title like"] = "%" + req.Title + "%"
	}
	if req.CategoryId > 0 {
		whereMap["category_id"] = req.CategoryId
	}
	build, vars, err := daos.WhereBuild(whereMap)
	if err != nil {
		return nil, err
	}
	if len(vars) > 0 {
		db = db.Where(build, vars)
	}
	s.dao.DB.Find(&tags)
	_ = db.Limit(req.GetPageSize()).Offset(req.GetOffset()).Order("updated_at DESC").Find(&articles).Limit(-1).Offset(-1).Count(&total)
	for _, tag := range tags {
		tagsMap[cast.ToInt(tag.Id)] = tag.Name
	}
	for _, article := range articles {
		var (
			tagIds   []uint
			tagNames []string
		)
		err = s.dao.DB.Where("article_id=?", article.Id).Find(&articleTags).Error
		if err != nil {
			logger.Error("not found article with tag")
		} else {
			for _, tag := range articleTags {
				tagIds = append(tagIds, tag.TagId)
				tagNames = append(tagNames, tagsMap[cast.ToInt(tag.TagId)])
			}
		}

		articleList = append(articleList, models.ArticleInfo{
			Id:          article.Id,
			Title:       article.Title,
			Content:     article.Content,
			Description: article.Desc,
			TagIds:      tagIds,
			TagNames:    tagNames,
			CreateDate:  cast.ToString(article.CreatedAt),
			UpdateDate:  cast.ToString(article.UpdatedAt),
		})
	}

	return models.FrontArticleListRes{
		PageInfo: models.PageInfo{
			Page:      req.GetPage(),
			Total:     total,
			PageSize:  req.GetPageSize(),
			TotalPage: req.GetTotalPage(total),
		},
		Data: articleList,
	}, nil
}

func (s *ArticleService) GetFrontDetail(ctx *gin.Context, id string) (interface{}, error) {
	var (
		article     models.Article
		category    models.Category
		articleTags []models.ArticleTag
		tags        []models.Tag
		tagIds      []uint
		tagNames    []string
		tagMap      = make(map[uint]string)
	)
	err := s.dao.DB.Where("id=?", id).Find(&article).Error
	if err != nil {
		return nil, err
	}
	if article.Id == 0 {
		return nil, errors.New("文章不存在")
	}
	err = s.dao.DB.Where("article_id=?", id).Find(&articleTags).Error
	if err != nil {
		return nil, err
	}
	err = s.dao.DB.Where("id=?", article.CategoryId).Find(&category).Error
	if err != nil {
		return nil, err
	}
	err = s.dao.DB.Find(&tags).Error
	if err != nil {
		return nil, err
	}
	for _, tag := range tags {
		tagMap[tag.Id] = tag.Name
	}

	for _, articleTag := range articleTags {
		tagIds = append(tagIds, articleTag.TagId)
		if _, ok := tagMap[articleTag.TagId]; ok {
			tagNames = append(tagNames, tagMap[articleTag.TagId])
		}
	}

	return models.ArticleInfo{
		Id:          article.Id,
		Title:       article.Title,
		Content:     article.Content,
		Category:    category.Name,
		Description: article.Desc,
		CreateDate:  cast.ToString(article.CreatedAt),
		UpdateDate:  cast.ToString(article.UpdatedAt),
		TagIds:      tagIds,
		TagNames:    tagNames,
		CategoryId:  article.CategoryId,
	}, nil

}
