package main

import (
	"student_data/Config"

	"fmt"
	"github.com/jinzhu/gorm"
	"student_data/Models"
	"student_data/Routers"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Student{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}
