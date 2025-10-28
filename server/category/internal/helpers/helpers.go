package helpers

import (
	"strings"
	"github.com/sirupsen/logrus"
)

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

func ToTitle(s string) string {
	words := strings.Fields(s)
	for i, w := range words {
		if len(w) > 0 {
			words[i] = strings.ToUpper(string(w[0])) + strings.ToLower(w[1:])
		}
	}
	return strings.Join(words, " ")
}

func ToSlug(s string) string {
    words := strings.ToLower(s)
    words = strings.ReplaceAll(words, " ", "-")

    return words
}