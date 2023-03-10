package models

type User struct {
	Id       uint   `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Username string `gorm:"column:username;type:varchar(200);NOT NULL" json:"username"`
	Password string `gorm:"column:password;type:varchar(128);NOT NULL" json:"password"`
	Salt     string `gorm:"column:salt;type:varchar(20);NOT NULL" json:"salt"`
	Model
}

func (m *User) TableName() string {
	return "user"
}
