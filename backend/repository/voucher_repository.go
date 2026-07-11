package repository

import (
	"airline-voucher/backend/db/sqlc"
	"airline-voucher/backend/models"
	"context"
	"database/sql"
	"errors"
)

type VoucherRepository interface {
	CheckExists(ctx context.Context, flightNumber, date string) (bool, error)
	InsertVoucher(ctx context.Context, voucher *models.Voucher) error
}

type voucherRepositoryImpl struct {
	q *sqlc.Queries
}

func NewVoucherRepository(db *sql.DB) VoucherRepository {
	return &voucherRepositoryImpl{
		q: sqlc.New(db),
	}
}

func (r *voucherRepositoryImpl) CheckExists(ctx context.Context, flightNumber, date string) (bool, error) {
	_, err := r.q.CheckExists(ctx, sqlc.CheckExistsParams{
		FlightNumber: flightNumber,
		FlightDate:   date,
	})

	// Jika tidak ada baris yang ditemukan
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	// Jika ada error lain
	if err != nil {
		return false, err
	}

	// Ditemukan
	return true, nil
}

func (r *voucherRepositoryImpl) InsertVoucher(ctx context.Context, v *models.Voucher) error {
	err := r.q.InsertVoucher(ctx, sqlc.InsertVoucherParams{
		CrewName:     v.CrewName,
		CrewID:       v.CrewID,
		FlightNumber: v.FlightNumber,
		FlightDate:   v.FlightDate,
		AircraftType: v.AircraftType,
		Seat1:        v.Seat1,
		Seat2:        v.Seat2,
		Seat3:        v.Seat3,
		CreatedAt:    v.CreatedAt,
	})

	return err
}
