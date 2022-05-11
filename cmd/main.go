package main

import (
	"github.com/devbyP/webboard/pkgs/http/web"
	models "github.com/devbyP/webboard/pkgs/storage/db-gorm"
)

func main() {
	gorm := models.ConnectDB()
	gorm.AutoMigrate(&models.User{}, &models.Post{}, &models.Rep{})
	mapperFunc := web.URLMapping
	web.StartServer(mapperFunc)
}
