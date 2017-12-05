package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

//模型定义
//修改表名

type Product struct {
	ID    uint
	Code  string
	Price uint
}

//修改默认表名
func (Product) TableName() string {
	return "product2"
}

type Email struct {
	ID         int
	Email      string
}

func main() {
	//db, err := gorm.Open("mysql", "root:root@tcp(localhost:3306)/gorm")
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/test?charset=utf8")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	
	//设置默认表名前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		fmt.Println("prefix_" + defaultTableName)
		return "prefix_" + defaultTableName
	}
	
	//自动生成表
	db.AutoMigrate(&Product{}, &Email{})
	
}
