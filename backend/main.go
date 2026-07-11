package main

import (
	"airline-voucher/backend/config"
	"airline-voucher/backend/di"
	"airline-voucher/backend/routes"
	"log"
)

func main() {
	// 1. Load Environment Configuration
	cfg := config.LoadConfig()

	// 2. Setup Database Connection
	db, err := config.InitDatabase(cfg.DBPath)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	// Pastikan koneksi DB ditutup saat aplikasi berhenti
	defer db.Close()

	// 3. Initialize Dependency Injection
	h, err := di.InitializeHandler(db)
	if err != nil {
		log.Fatalf("Failed to initialize dependencies: %v", err)
	}

	// 4. Setup Echo Server
	e := config.InitEcho()

	// 5. Register Routes
	routes.SetupRoutes(e, h)

	// 6. Start Server
	log.Printf("Server berjalan di port :%s", cfg.Port)
	e.Logger.Fatal(e.Start(":" + cfg.Port))
}
