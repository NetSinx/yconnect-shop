package utils

import "log"

func PanicError(err error) {
	log.Panic(err)
}