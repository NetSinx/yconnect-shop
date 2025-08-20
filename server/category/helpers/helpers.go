package helpers

import "strings"

func ToTitle(s string) string {
    return strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
}