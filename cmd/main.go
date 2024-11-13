package main

import (
	"backend/config"
	"backend/internal/handlers"
	"backend/internal/models"
	"backend/internal/repositories"
	"backend/internal/router"
	"fmt"
	"log"

	_ "backend/docs" // This line is necessary for go-swagger to find docs

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//	@title			MyApp API
//	@version		1.0
//	@description	This is a sample server.

//	@host		localhost:8080
//	@BasePath	/

//	@schemes	http

func main() {
	cfg := config.NewConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Migrate the schema
	err = db.AutoMigrate(&models.Service{}, &models.Group{}, &models.ClassifiedService{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	serviceRepo := repositories.NewServiceRepository(db)
	groupRepo := repositories.NewGroupRepository(db)
	csRepo := repositories.NewClassifiedServiceRepository(db)
	handler := handlers.NewHandler(serviceRepo, groupRepo, csRepo)

	r := router.SetupRouter(handler)

	log.Fatal(r.Run(":8080"))
}
