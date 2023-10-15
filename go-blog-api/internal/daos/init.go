package daos

import (
	"github.com/convee/go-blog-api/pkg/storage/orm"
	"gorm.io/gorm"
)

// DB 数据库全局变量
var DB *gorm.DB

// Init 初始化数据库
func Init(cfg *orm.Config) *gorm.DB {
	DB = orm.NewMySQL(cfg)
	return DB
}

// GetDB 返回默认的数据库
func GetDB() *gorm.DB {
	return DB
}
