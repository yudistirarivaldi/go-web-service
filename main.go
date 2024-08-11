package main

import (
	"log"
	"web-service/config"
	_ "web-service/docs"
	"web-service/model"
	"web-service/routes"
	"web-service/seeders"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()

	db := config.InitDB()
	defer db.Close()

	model.DB = db

	// Jalankan migrasi
	if err := config.RunMigrations(db); err != nil {
		log.Fatalf("Error running migrations: %v", err)
	}

	// Jalankan seederds
	if err := seeders.SeedUsers(db); err != nil {
		log.Fatalf("Error seeding data: %v", err)
	}

	routes.UserRoutes(r)

	 r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	

	r.Run(":8080")
}
