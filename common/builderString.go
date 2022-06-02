package common

import "strings"

func BuilderString(a string, b string) string {
	var sb = new(strings.Builder)
	sb.Grow(len(b) + len(a))

	sb.WriteString(a)
	sb.WriteString(b)

	return sb.String()
}
