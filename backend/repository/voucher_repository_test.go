package repository_test

import (
	"airline-voucher/backend/models"
	"airline-voucher/backend/repository"
	"context"
	"database/sql"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func setupTestDB(t *testing.T) *sql.DB {
	// Gunakan SQLite in-memory untuk testing yang cepat dan terisolasi
	db, err := sql.Open("sqlite3", "file::memory:?cache=shared")
	assert.NoError(t, err)

	// Buat tabel manual untuk keperluan test
	query := `
	CREATE TABLE vouchers (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		crew_name TEXT NOT NULL,
		crew_id TEXT NOT NULL,
		flight_number TEXT NOT NULL,
		flight_date TEXT NOT NULL,
		aircraft_type TEXT NOT NULL,
		seat1 TEXT NOT NULL,
		seat2 TEXT NOT NULL,
		seat3 TEXT NOT NULL,
		created_at TEXT NOT NULL
	);`
	_, err = db.Exec(query)
	assert.NoError(t, err)

	return db
}

func TestVoucherRepository(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := repository.NewVoucherRepository(db)

	t.Run("Insert and CheckExists Success", func(t *testing.T) {
		voucher := &models.Voucher{
			CrewName:     "Budi",
			CrewID:       "12345",
			FlightNumber: "GA101",
			FlightDate:   "2025-10-10",
			AircraftType: "ATR",
			Seat1:        "1A",
			Seat2:        "2B",
			Seat3:        "3C",
			CreatedAt:    time.Now().Format(time.RFC3339),
		}

		// Test Insert
		err := repo.InsertVoucher(context.Background(), voucher)
		assert.NoError(t, err)

		// Test CheckExists (Harus true)
		exists, err := repo.CheckExists(context.Background(), "GA101", "2025-10-10")
		assert.NoError(t, err)
		assert.True(t, exists)
	})

	t.Run("CheckExists NotFound", func(t *testing.T) {
		// Harusnya tidak ada penerbangan GA999
		exists, err := repo.CheckExists(context.Background(), "GA999", "2025-10-10")
		assert.NoError(t, err)
		assert.False(t, exists)
	})
}
