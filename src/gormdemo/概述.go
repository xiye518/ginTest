package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Product struct {
	ID uint `gorm:"primary_key"`
	Code string
	Price uint
}

//概述
func main() {
	//db,err := gorm.Open("mysql","root:123456@tcp(localhost:3306)/gorm")
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/test?charset=utf8")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	
	// 自动迁移表，生成的表名为 products
	db.AutoMigrate(&Product{})
	
	// Create
	db.Create(&Product{Code: "L1212", Price: 1000})
	
	// Read
	var product Product
	db.First(&product, 1) // find product with id 1
	db.First(&product, "code = ?", "L1212") // find product with code l1212
	
	// Update
	db.Model(&product).Update("Price", 2000)
	
	//Delete
	//db.Delete(&product)
	
}
