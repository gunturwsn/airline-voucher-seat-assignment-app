package voucher

import (
	"backend/models"
)

type VoucherService interface {
	CheckVoucherExists(flightNumber, date string) (bool, error)
	GenerateVoucher(data models.Voucher) ([]string, error)
}
