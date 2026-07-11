package handler

import (
	"airline-voucher/backend/models"
	"airline-voucher/backend/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type VoucherHandler struct {
	service usecase.VoucherService
}

func NewVoucherHandler(s usecase.VoucherService) *VoucherHandler {
	return &VoucherHandler{service: s}
}

func (h *VoucherHandler) Check(c echo.Context) error {
	var req models.CheckRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	exists, err := h.service.CheckVoucher(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, models.CheckResponse{Exists: exists})
}

func (h *VoucherHandler) Generate(c echo.Context) error {
	var req models.GenerateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, models.GenerateResponse{Success: false, Error: "Invalid request body"})
	}

	seats, err := h.service.GenerateVoucher(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.GenerateResponse{Success: false, Error: err.Error()})
	}

	return c.JSON(http.StatusOK, models.GenerateResponse{
		Success: true,
		Seats:   seats,
	})
}
