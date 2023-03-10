package models

type Page struct {
	Id      uint   `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Ident   string `gorm:"column:ident;type:varchar(20);NOT NULL" json:"ident"`
	Title   string `gorm:"column:title;type:varchar(100);NOT NULL" json:"title"`
	Content string `gorm:"column:content;type:longtext" json:"content"`
	Model
}

func (m *Page) TableName() string {
	return "page"
}

type PageListReq struct {
	PageInfo
	Keyword string `json:"keyword"`
}

type PageListRes struct {
	PageInfo
	Data []Page `json:"data"`
}

type PageAddReq struct {
	Title string `json:"title" validate:"required"`
}

type PageUpdateReq struct {
	Id    uint   `json:"id"  validate:"required"`
	Title string `json:"title" validate:"required"`
}

type PageDelReq struct {
	Id uint `json:"id" validate:"required"`
}
