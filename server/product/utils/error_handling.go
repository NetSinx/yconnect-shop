package utils

import "log"

func PanicError(err error) {
	log.Panic(err)
}

func failOnError(msg string, err error) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}