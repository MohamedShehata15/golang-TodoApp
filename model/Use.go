package model

type User struct {
	Id       uint   `gorm:"primary_key; auto_increment" json:"id"`
	Name     string `gorm:"varchar(255); not null" json:"name"`
	Email    string `gorm:"varchar(255); not null; unique" json:"email"`
	Password string `gorm:"->;<-; not null" json:"-"`
	Token    string `gorm:"-" json:"token,omitempty"`
}