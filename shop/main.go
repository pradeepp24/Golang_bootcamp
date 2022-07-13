package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"retail_shop/Config"
	"retail_shop/Models"
	"retail_shop/Routers"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status", err)
	}
	Config.DB.AutoMigrate(&Models.Customer{}, &Models.Order{}, &Models.Product{})
	r := Routes.SetupRouter()
	r.Run()
}
