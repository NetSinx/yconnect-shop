package test

import (
	"strings"
	"testing"
)

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