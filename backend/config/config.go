package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port   string
	DBPath string
}

func LoadConfig() *Config {
	// Abaikan error jika .env tidak ada (berguna jika jalan di Docker via ENV vars)
	_ = godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default fallback
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "./vouchers.db"
	}

	return &Config{
		Port:   port,
		DBPath: dbPath,
	}
}
