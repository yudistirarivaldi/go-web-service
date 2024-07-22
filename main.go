package main

import (
	"web-service/config"
	"web-service/model"
	"web-service/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db := config.InitDB()
	defer db.Close() 

	model.DB = db

	routes.UserRoutes(r)

	r.Run(":8080")
}
