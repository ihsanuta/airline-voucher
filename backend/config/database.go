package config

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

func InitDatabase(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	// Jalankan Migrasi Otomatis
	log.Println("Menjalankan migrasi database...")
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return nil, err
	}

	dir, _ := os.Getwd()
	// Gabungkan dengan path migrasi
	migrationPath := "file://" + filepath.Join(dir, "db", "migrations")

	// Menunjuk ke folder db/migration
	m, err := migrate.NewWithDatabaseInstance(
		migrationPath,
		"sqlite3", driver)
	if err != nil {
		return nil, err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return nil, err
	}
	log.Println("Migrasi database berhasil.")

	return db, nil
}
