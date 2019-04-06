package main

import (
	"fmt"
	"go-rest-api/app/controller"
	"go-rest-api/app/model"
	"go-rest-api/app/route"
	"go-rest-api/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	db := config.ConnectToDb()

	controller.Database = db
	model.Database = db
}

func main() {
	r := route.Init()

	fmt.Println(r.Run(":3000"))
}
