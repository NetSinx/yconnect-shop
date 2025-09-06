package helpers

import "github.com/sirupsen/logrus"

func FatalError(log *logrus.Logger, err error, msg string) {
  if err != nil {
    log.Fatalf("%s: %s", msg, err)
  }
}

func PanicError(log *logrus.Logger, err error, msg string) {
  if err != nil {
    log.Panicf("%s: %s", msg, err)
  }
}