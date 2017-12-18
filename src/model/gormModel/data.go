package gormModel

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//var db *gorm.DB

//func init() {
//	var err error
//	db, err = gorm.Open("mysql", "root:123456@tcp(localhost:3306)/gorm?charset=utf8")
//	if err != nil {
//		panic("failed to connect database")
//	}
//
//}

func CreatTable() *gorm.DB {
	db_gorm, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/gorm?charset=utf8")
	if err != nil {
		panic("failed to connect database")
	}
	defer db_gorm.Close()
	
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "Profile_" + defaultTableName
	}
	
	creat := db_gorm.AutoMigrate(&Product{}, &Email{})
	return creat
}
