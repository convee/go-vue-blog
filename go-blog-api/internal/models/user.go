package models

type User struct {
	Id       uint   `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Name     string `gorm:"column:name;type:varchar(200);NOT NULL" json:"name"`
	Password string `gorm:"column:password;type:varchar(128);NOT NULL" json:"password"`
	Avatar   string `gorm:"column:avatar;type:text" json:"avatar"`
	Salt     string `gorm:"column:salt;type:varchar(20);NOT NULL" json:"salt"`
	Model
}

func (m *User) TableName() string {
	return "user"
}
