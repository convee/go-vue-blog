package models

type Article struct {
	Id         uint   `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Title      string `gorm:"column:title;type:varchar(200);NOT NULL" json:"title"`
	Status     int    `gorm:"column:status;type:tinyint(1);default:0;NOT NULL" json:"status"`
	ImgUrl     string `gorm:"column:img_url;type:varchar(2000);NOT NULL" json:"img_url"`
	CategoryId int    `gorm:"column:category_id;type:int(11);default:0;NOT NULL" json:"category_id"`
	IsTop      int    `gorm:"column:is_top;type:tinyint(1);default:0;NOT NULL" json:"is_top"`
	Views      int    `gorm:"column:views;type:int(11);default:0;NOT NULL" json:"views"`
	Desc       string `gorm:"column:desc;type:varchar(2000);NOT NULL" json:"desc"`
	Content    string `gorm:"column:content;type:longtext" json:"content"`
	Model
}

func (m *Article) TableName() string {
	return "article"
}

type ArticleListReq struct {
	PageInfo
	Keyword    string `json:"keyword"`
	CategoryId int    `json:"category_id"`
}

type ArticleListRes struct {
	PageInfo
	Data []Article `json:"data"`
}

type ArticleInfoReq struct {
}
