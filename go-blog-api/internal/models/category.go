package models

type Category struct {
	Id   uint   `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Name string `gorm:"column:name;type:varchar(100);NOT NULL" json:"name"`
	Model
}

func (m *Category) TableName() string {
	return "category"
}

type CategoryListReq struct {
	PageInfo
	Name string `json:"name" form:"name"`
}

type CategoryListRes struct {
	PageInfo
	PageData []Category `json:"pageData"`
}

type CategoryAddReq struct {
	Name string `json:"name" validate:"required"`
}

type CategoryUpdateReq struct {
	Id   uint   `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type CategoryDelReq struct {
	Id uint `json:"id" validate:"required"`
}
