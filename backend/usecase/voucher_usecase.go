package usecase

import (
	"airline-voucher/backend/models"
	"airline-voucher/backend/repository"
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type VoucherService interface {
	CheckVoucher(ctx context.Context, req models.CheckRequest) (bool, error)
	GenerateVoucher(ctx context.Context, req models.GenerateRequest) ([]string, error)
}

type voucherServiceImpl struct {
	repo repository.VoucherRepository
}

func NewVoucherService(repo repository.VoucherRepository) VoucherService {
	return &voucherServiceImpl{repo: repo}
}

func (s *voucherServiceImpl) CheckVoucher(ctx context.Context, req models.CheckRequest) (bool, error) {
	if req.FlightNumber == "" || req.Date == "" {
		return false, errors.New("flightNumber and date are required")
	}
	return s.repo.CheckExists(ctx, req.FlightNumber, req.Date)
}

func (s *voucherServiceImpl) GenerateVoucher(ctx context.Context, req models.GenerateRequest) ([]string, error) {
	// 1. Validasi input
	if req.Name == "" || req.ID == "" || req.FlightNumber == "" || req.Date == "" || req.Aircraft == "" {
		return nil, errors.New("all fields are required")
	}

	// 2. Generate valid seats
	seats, err := generateUniqueSeats(req.Aircraft)
	if err != nil {
		return nil, err
	}

	// 3. Simpan ke database
	voucher := &models.Voucher{
		CrewName:     req.Name,
		CrewID:       req.ID,
		FlightNumber: req.FlightNumber,
		FlightDate:   req.Date,
		AircraftType: req.Aircraft,
		Seat1:        seats[0],
		Seat2:        seats[1],
		Seat3:        seats[2],
		CreatedAt:    time.Now().UTC().Format(time.RFC3339),
	}

	if err := s.repo.InsertVoucher(ctx, voucher); err != nil {
		return nil, errors.New("failed to save voucher to database")
	}

	return seats, nil
}

// Algoritma pembuatan 3 kursi unik berdasarkan jenis pesawat
func generateUniqueSeats(aircraft string) ([]string, error) {
	var maxRow int
	var cols []string

	switch aircraft {
	case "ATR":
		maxRow = 18
		cols = []string{"A", "C", "D", "F"}
	case "Airbus 320", "Boeing 737 Max":
		maxRow = 32
		cols = []string{"A", "B", "C", "D", "E", "F"}
	default:
		return nil, fmt.Errorf("invalid aircraft type: %s", aircraft)
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	seatSet := make(map[string]bool)
	var seats []string

	for len(seats) < 3 {
		row := r.Intn(maxRow) + 1
		col := cols[r.Intn(len(cols))]
		seat := fmt.Sprintf("%d%s", row, col)

		if !seatSet[seat] {
			seatSet[seat] = true
			seats = append(seats, seat)
		}
	}
	return seats, nil
}
