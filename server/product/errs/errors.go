package errs

import "log"

var (
	ErrDuplicatedKey = "Produk sudah tersedia."
	ErrNotFound = "Produk tidak ditemukan."
	ErrInternalServer = "Terjadi masalah pada server."
)

func PanicError(err error) {
	log.Panic(err)
}

func FailOnError(msg string, err error) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}