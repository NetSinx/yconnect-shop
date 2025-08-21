package helpers

import "strings"

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