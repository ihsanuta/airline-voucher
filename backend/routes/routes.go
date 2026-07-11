package routes

import (
	"airline-voucher/backend/handler"

	"github.com/labstack/echo/v4"
)

// SetupRoutes menerima instance Echo dan menghubungkannya dengan handler
func SetupRoutes(e *echo.Echo, voucherHandler *handler.VoucherHandler) {
	// Grouping untuk versi API
	api := e.Group("/api")
	{
		api.POST("/check", voucherHandler.Check)
		api.POST("/generate", voucherHandler.Generate)
	}
}
