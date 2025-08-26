package helpers

import (
	"strings"
	"github.com/sirupsen/logrus"
)

type Helpers struct {
	Log *logrus.Logger
}

func NewHelpers(log *logrus.Logger) *Helpers {
	return &Helpers{
		Log: log,
	}
}

func (h *Helpers) FatalError(err error, msg string) {
  if err != nil {
    h.Log.Fatalf("%s: %s", msg, err)
  }
}

func (h *Helpers) PanicError(err error, msg string) {
  if err != nil {
    h.Log.Panicf("%s: %s", msg, err)
  }
}

func (h *Helpers) ToTitle(s string) string {
	words := strings.Fields(s)
	for i, w := range words {
		if len(w) > 0 {
			words[i] = strings.ToUpper(string(w[0])) + strings.ToLower(w[1:])
		}
	}
	return strings.Join(words, " ")
}

func (h *Helpers) ToSlug(s string) string {
    words := strings.ToLower(s)
    words = strings.ReplaceAll(words, " ", "-")

    return words
}