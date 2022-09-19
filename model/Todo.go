package model

type Todo struct {
	Id        uint   `gorm:"primary_key; auto_increment" json:"id"`
	Title     string `gorm:"varchar(255); not null" json:"title"`
	Completed bool   `gorm:"default:false" json:"completed"`
	UserId    uint   `gorm:"not null" json:"-"`
	User      User   `gorm:"foreignKey:UserId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}