package main

import (
	"github.com/devbyP/webboard/models"
	"github.com/devbyP/webboard/server"
)

func main() {
	gorm := models.ConnectDB()
	gorm.AutoMigrate(&models.User{}, &models.Post{}, &models.Rep{})
	mapperFunc := server.URLMapping
	server.StartServer(mapperFunc)
}
