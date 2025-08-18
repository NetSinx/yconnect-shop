package utils

import (
	"crypto/rand"
	"encoding/base64"
	"strings"
)

func GenerateSlugByName(name string) string {
	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, " ", "-")

	b := make([]byte, 16)
	rand.Read(b)
	uniqueId := base64.URLEncoding.EncodeToString(b)

	return name + "-" + uniqueId
}