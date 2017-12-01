package gormModel

import (
	"time"
	"github.com/jinzhu/gorm"
)

//设置字段

type Product struct {
	ID    uint `gorm:"primary_key:id"`
	Num   int  `gorm:"AUTO_INCREMENT:number"`
	Code  string
	Price uint
	Tag   []Tag     `gorm:"many2many:tag;"`
	Date  time.Time `gorm:"-"`
}

type Email struct {
	ID         int    `gorm:"primary_key:id"`
	UserID     int    `gorm:"not null;index"`
	Email      string `gorm:"type:varchar(100);unique_index"`
	Subscribed bool
}

type Tag struct {
	Name string
}

//设置外键字段

type Profile struct {
	gorm.Model
	Refer int
	Name  string
}

type UserInfo struct {
	gorm.Model
	Profile   Profile `gorm:"ForeignKey:ProfileID;AssociationForeignKey:Refer"`
	ProfileID int
}

