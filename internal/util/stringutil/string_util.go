package stringutil

import "strings"

func BuildStringWithParams(str ...string) string {
	var sb strings.Builder
	for _, value := range str {
		sb.WriteString(value)
	}

	return sb.String()
}
