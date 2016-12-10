package main

import (
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)


type Product struct {
    gorm.Model
    Code string
    Price uint
}

func main()  {
    db, err := gorm.Open("mysql", "root:@/gormdemo?charset=utf8&parseTime=True&loc=Local")
    if err != nil {
        panic("failed to connect database")
    }

    defer db.Close()

    db.AutoMigrate(&Product{})

    db.Create(&Product{Code:"L1212", Price:1000})

    var product Product
    db.First(&product, 1)
    db.First(&product, "code = ?", "L1212")
    fmt.Print(product)

    db.Model(&product).Update("Price", 2000)
    
}