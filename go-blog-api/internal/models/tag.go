package models

type Tag struct {
	Id   uint   `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Name string `gorm:"column:name;type:varchar(100);NOT NULL" json:"name"`
	Model
}

func (m *Tag) TableName() string {
	return "tag"
}

type TagAddReq struct {
	Name string `json:"name" validate:"required"`
}

type TagUpdateReq struct {
	Id   uint   `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type TagDelReq struct {
	Id uint `json:"id" validate:"required"`
}

type ArticleTag struct {
	Id        uint `gorm:"column:id;type:int(10) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	ArticleId int  `gorm:"column:article_id;type:int(11);comment:文章 ID;NOT NULL" json:"article_id"`
	TagId     uint `gorm:"column:tag_id;type:int(10) unsigned;default:0;comment:标签 ID;NOT NULL" json:"tag_id"`
}

func (m *ArticleTag) TableName() string {
	return "article_tag"
}

type TagListReq struct {
	PageInfo
	Name string `json:"name" form:"name"`
}

type TagListRes struct {
	PageInfo
	PageData []Tag `json:"pageData"`
}
