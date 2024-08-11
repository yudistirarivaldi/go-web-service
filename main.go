package main

import (
	"log"
	"web-service/config"
	"web-service/model"
	"web-service/routes"
	"web-service/seeders"

	"github.com/gin-gonic/gin"
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

	r.Run(":8080")
}
