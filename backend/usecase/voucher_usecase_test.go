package usecase_test

import (
	"airline-voucher/backend/models"
	"airline-voucher/backend/usecase"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// --- MOCK REPOSITORY ---
type MockVoucherRepository struct {
	mock.Mock
}

func (m *MockVoucherRepository) CheckExists(ctx context.Context, flightNumber, date string) (bool, error) {
	args := m.Called(flightNumber, date)
	return args.Bool(0), args.Error(1)
}

func (m *MockVoucherRepository) InsertVoucher(ctx context.Context, voucher *models.Voucher) error {
	args := m.Called(voucher)
	return args.Error(0)
}

// --- TESTS ---
func TestService_CheckVoucher(t *testing.T) {
	mockRepo := new(MockVoucherRepository)
	svc := usecase.NewVoucherService(mockRepo)

	t.Run("Success", func(t *testing.T) {
		// Definisikan ekspektasi mock
		mockRepo.On("CheckExists", "GA102", "2025-07-12").Return(true, nil).Once()

		req := models.CheckRequest{FlightNumber: "GA102", Date: "2025-07-12"}
		exists, err := svc.CheckVoucher(context.Background(), req)

		assert.NoError(t, err)
		assert.True(t, exists)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Validation Error", func(t *testing.T) {
		req := models.CheckRequest{FlightNumber: "", Date: ""}
		exists, err := svc.CheckVoucher(context.Background(), req)

		assert.Error(t, err)
		assert.Equal(t, "flightNumber and date are required", err.Error())
		assert.False(t, exists)
	})
}

func TestService_GenerateVoucher(t *testing.T) {
	mockRepo := new(MockVoucherRepository)
	svc := usecase.NewVoucherService(mockRepo)

	t.Run("Success Generate Seats", func(t *testing.T) {
		// Mock InsertVoucher agar selalu me-return nil (tanpa error)
		mockRepo.On("InsertVoucher", mock.AnythingOfType("*models.Voucher")).Return(nil).Once()

		req := models.GenerateRequest{
			Name:         "Sarah",
			ID:           "98123",
			FlightNumber: "ID102",
			Date:         "2025-07-12",
			Aircraft:     "Airbus 320",
		}

		seats, err := svc.GenerateVoucher(context.Background(), req)

		assert.NoError(t, err)
		assert.Len(t, seats, 3) // Harus meng-generate tepat 3 kursi
		mockRepo.AssertExpectations(t)
	})

	t.Run("Invalid Aircraft", func(t *testing.T) {
		req := models.GenerateRequest{
			Name:         "John",
			ID:           "111",
			FlightNumber: "XX123",
			Date:         "2025-01-01",
			Aircraft:     "UFO", // Invalid aircraft
		}

		seats, err := svc.GenerateVoucher(context.Background(), req)

		assert.Error(t, err)
		assert.Nil(t, seats)
		assert.Contains(t, err.Error(), "invalid aircraft type")
	})
}
