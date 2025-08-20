package test

import (
	"crypto/rand"
	"encoding/base64"
	"strings"
	"testing"
)

func TestGenerateSlugByName(t *testing.T) {
	name := "Ayam Bakar Pedas Manis"

	trimmed := strings.TrimSpace(name)
	words := strings.Fields(trimmed)
	name = strings.Join(words, " ")
	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, " ", "-")

	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		t.Errorf("Error: %v", err)
	}

	uid := base64.URLEncoding.EncodeToString(b)
	slug := name + "-" + uid

	t.Logf("Result: %v", slug)
}

func TestReplaceProductSlug(t *testing.T) {
	nama := "Ayam Bakar Pedas Manis"
	
	slug := "ayam-bakar-ergjklgr34gg=="
	expected := "ayam-bakar-pedas-manis-ergjklgr34gg=="
	splitSlug := strings.Split(slug, "-")
	uid := splitSlug[len(splitSlug) - 1]

	trimmed := strings.TrimSpace(nama)
	words := strings.Fields(trimmed)
	nama = strings.Join(words, " ")
	nama = strings.ToLower(nama)
	nama = strings.ReplaceAll(nama, " ", "-")
	newSlug := nama + "-" + uid

	if newSlug != expected {
		t.Errorf("Slug doesn't match with expected. Result: %s", newSlug)
	}
}