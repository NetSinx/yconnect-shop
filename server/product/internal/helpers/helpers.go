package helpers

import (
	"crypto/rand"
	"encoding/base64"
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

func GenerateSlugByName(name string) (string, error) {
	trimmed := strings.TrimSpace(name)
	words := strings.Fields(trimmed)
	name = strings.Join(words, " ")
	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, " ", "-")

	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	uniqueId := base64.URLEncoding.EncodeToString(b)
	slug := name + "-" + uniqueId

	return slug, nil
}

func ReplaceProductSlug(slug string, namaProduct string) string {
	splitSlug := strings.Split(slug, "-")
	uid := splitSlug[len(splitSlug) - 1]

	trimmed := strings.TrimSpace(namaProduct)
	words := strings.Fields(trimmed)
	namaProduct = strings.Join(words, " ")
	namaProduct = strings.ToLower(namaProduct)
	namaProduct = strings.ReplaceAll(namaProduct, " ", "-")
	newSlug := namaProduct + "-" + uid

	return newSlug
}