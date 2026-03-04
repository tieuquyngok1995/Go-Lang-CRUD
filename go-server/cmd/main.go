package main

import (
	"go-crud/internal/config"
	"go-crud/internal/database"
	"go-crud/internal/router"
	"log"
)

func main() {
	cfg := config.Load()
	db := database.InitDB(cfg.DB)
	defer db.Close()

	r := router.SetupUserRoutes(db)

	log.Printf("? Server running on port %s", cfg.AppPort)
	if err := r.Run(":" + cfg.AppPort); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
