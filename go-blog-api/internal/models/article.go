package models

type Article struct {
	Id         uint   `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Title      string `gorm:"column:title;type:varchar(200);NOT NULL" json:"title"`
	Status     int    `gorm:"column:status;type:tinyint(1);default:0;NOT NULL" json:"status"`
	ImgUrl     string `gorm:"column:img_url;type:varchar(2000);NOT NULL" json:"imgUrl"`
	CategoryId int    `gorm:"column:category_id;type:int(11);default:0;NOT NULL" json:"categoryId"`
	IsTop      int    `gorm:"column:is_top;type:tinyint(1);default:0;NOT NULL" json:"isTop"`
	Views      int    `gorm:"column:views;type:int(11);default:0;NOT NULL" json:"views"`
	Desc       string `gorm:"column:desc;type:varchar(2000);NOT NULL" json:"desc"`
	Content    string `gorm:"column:content;type:longtext" json:"content"`
	Model
}

type ArticleInfo struct {
	Id          uint     `json:"id"`
	Title       string   `json:"title"`
	Content     string   `json:"content"`
	Category    string   `json:"category"`
	Description string   `json:"description"`
	Author      string   `json:"author"`
	CreateDate  string   `json:"createDate"`
	UpdateDate  string   `json:"updateDate"`
	TagIds      []uint   `json:"tagIds"`
	TagNames    []string `json:"tagNames"`
	CategoryId  int      `json:"categoryId"`
}

func (m *Article) TableName() string {
	return "article"
}

type StatRes struct {
	ArticleCount  int64 `json:"articleCount"`
	CategoryCount int64 `json:"categoryCount"`
	PageCount     int64 `json:"pageCount"`
	TagCount      int64 `json:"tagCount"`
}
type ArticleListReq struct {
	PageInfo
	Title      string `json:"title" form:"title"`
	CategoryId int    `json:"categoryId" form:"categoryId"`
}

type ArticleListRes struct {
	PageInfo
	PageData []ArticleInfo `json:"pageData"`
}

type ArticleInfoReq struct {
}

type ArticleAddReq struct {
	Title      string `json:"title" validate:"required"`
	Content    string `json:"content" validate:"required"`
	CategoryId int    `json:"categoryId" validate:"required"`
	TagIds     []uint `json:"tagIds" validate:"required"`
}

type ArticleUpdateReq struct {
	Id         uint   `json:"id" validate:"required"`
	Title      string `json:"title" validate:"required"`
	Content    string `json:"content" validate:"required"`
	CategoryId int    `json:"categoryId" validate:"required"`
	TagIds     []uint `json:"tagIds" validate:"required"`
}

type ArticleDelReq struct {
	Id uint `json:"id" validate:"required"`
}

type FrontArticleListReq struct {
	PageInfo
	Title      string `json:"title" form:"title"`
	CategoryId int    `json:"categoryId" form:"categoryId"`
}
type FrontArticleListRes struct {
	PageInfo
	Data []Article `json:"data"`
}
