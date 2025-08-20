package errs

import "log"

var (
	ErrDuplicatedKey = "Produk sudah tersedia."
	ErrNotFound = "Produk tidak ditemukan."
	ErrInternalServer = "Terjadi masalah pada server."
)

func PanicError(err error) {
	if err != nil {
		log.Panic(err)
	}
}