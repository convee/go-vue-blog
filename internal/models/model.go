package models

import (
	"github.com/convee/go-vue-blog/internal/pkg/common"
	"gorm.io/gorm"
)

type Model struct {
	CreatedAt common.JSONTime `json:"created_at"`
	UpdatedAt common.JSONTime `json:"updated_at"`
	DeletedAt gorm.DeletedAt  `json:"deleted_at"`
}

type PageInfo struct {
	Page    int   `form:"page" json:"page"`
	Total   int64 `form:"total" json:"total"`
	PerPage int   `form:"per_page" json:"per_page"`
}

type OrderSort struct {
	OrderBy string `form:"order_by" json:"order_by,omitempty"`
	Sort    string `form:"sort" json:"sort,omitempty"`
}

func (p PageInfo) GetLimit() int {
	if p.PerPage <= 0 {
		return 20
	}
	return p.PerPage
}

func (p PageInfo) GetPage() int {
	if p.Page <= 0 {
		return 1
	}
	return p.Page
}

func (p PageInfo) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}
