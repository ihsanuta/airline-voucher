//go:build wireinject
// +build wireinject

package di

import (
	"airline-voucher/backend/handler"
	"airline-voucher/backend/repository"
	"airline-voucher/backend/usecase"
	"database/sql"

	"github.com/google/wire"
)

func InitializeHandler(db *sql.DB) (*handler.VoucherHandler, error) {
	wire.Build(
		repository.NewVoucherRepository,
		usecase.NewVoucherService,
		handler.NewVoucherHandler,
	)
	return &handler.VoucherHandler{}, nil
}
