package models

import (
	"github.com/convee/go-blog-api/internal/pkg/common"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"math"
)

type Model struct {
	CreatedAt common.JSONTime `json:"createdAt"`
	UpdatedAt common.JSONTime `json:"updatedAt"`
	DeletedAt gorm.DeletedAt  `json:"deletedAt"`
}

type PageInfo struct {
	Page      int   `form:"pageNo" json:"pageNum"`
	Total     int64 `form:"total" json:"total"`
	PageSize  int   `form:"pageSize" json:"pageSize"`
	TotalPage int64 `json:"totalPage"`
}

type OrderSort struct {
	OrderBy string `form:"orderBy" json:"orderBy,omitempty"`
	Sort    string `form:"sort" json:"sort,omitempty"`
}

func (p PageInfo) GetPageSize() int {
	if p.PageSize <= 0 {
		return 20
	}
	return p.PageSize
}

func (p PageInfo) GetPage() int {
	if p.Page <= 0 {
		return 1
	}
	return p.Page
}

func (p PageInfo) GetOffset() int {
	return (p.GetPage() - 1) * p.GetPageSize()
}

func (p PageInfo) GetTotalPage(total int64) int64 {
	ceil := math.Ceil(float64(total) / float64(p.GetPageSize()))
	return cast.ToInt64(ceil)
}
