package errs

import "log"

func PanicError(err error) {
	log.Panic(err)
}

func FailOnError(msg string, err error) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}