package handler_test

import (
	"airline-voucher/backend/handler"
	"airline-voucher/backend/models"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// --- MOCK USECASE ---
type MockVoucherUsecase struct {
	mock.Mock
}

func (m *MockVoucherUsecase) CheckVoucher(ctx context.Context, req models.CheckRequest) (bool, error) {
	args := m.Called(req)
	return args.Bool(0), args.Error(1)
}

func (m *MockVoucherUsecase) GenerateVoucher(ctx context.Context, req models.GenerateRequest) ([]string, error) {
	args := m.Called(req)
	var seats []string
	if args.Get(0) != nil {
		seats = args.Get(0).([]string)
	}
	return seats, args.Error(1)
}

// --- TESTS ---
func TestHandler_Check(t *testing.T) {
	mockSvc := new(MockVoucherUsecase)
	h := handler.NewVoucherHandler(mockSvc)
	e := echo.New()

	t.Run("Success", func(t *testing.T) {
		reqBody := models.CheckRequest{FlightNumber: "GA102", Date: "2025-07-12"}
		mockSvc.On("CheckVoucher", reqBody).Return(true, nil).Once()

		bodyJSON, _ := json.Marshal(reqBody)
		req := httptest.NewRequest(http.MethodPost, "/api/check", bytes.NewBuffer(bodyJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := h.Check(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		var resp models.CheckResponse
		json.Unmarshal(rec.Body.Bytes(), &resp)
		assert.True(t, resp.Exists)
		mockSvc.AssertExpectations(t)
	})
}

func TestHandler_Generate(t *testing.T) {
	mockSvc := new(MockVoucherUsecase)
	h := handler.NewVoucherHandler(mockSvc)
	e := echo.New()

	t.Run("Success", func(t *testing.T) {
		reqBody := models.GenerateRequest{
			Name: "Sarah", ID: "98123", FlightNumber: "ID102",
			Date: "2025-07-12", Aircraft: "Airbus 320",
		}
		expectedSeats := []string{"3B", "7C", "14D"}

		mockSvc.On("GenerateVoucher", reqBody).Return(expectedSeats, nil).Once()

		bodyJSON, _ := json.Marshal(reqBody)
		req := httptest.NewRequest(http.MethodPost, "/api/generate", bytes.NewBuffer(bodyJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := h.Generate(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		var resp models.GenerateResponse
		json.Unmarshal(rec.Body.Bytes(), &resp)

		assert.True(t, resp.Success)
		assert.ElementsMatch(t, expectedSeats, resp.Seats)
		mockSvc.AssertExpectations(t)
	})
}
