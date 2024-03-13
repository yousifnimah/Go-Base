package Helper

import (
	"regexp"
	"strings"
)

func EscapeSChars(s string) string {
	chars := []string{"]", "^", "\\\\", "[", ".", "(", ")", "-"}
	r := strings.Join(chars, "")
	re := regexp.MustCompile("[" + r + "]+")
	s = re.ReplaceAllString(s, "")
	s = trimQuotes(s)
	return s
}

func trimQuotes(s string) string {
	if len(s) >= 2 {
		if c := s[len(s)-1]; s[0] == c && (c == '"' || c == '\'') {
			return s[1 : len(s)-1]
		}
	}
	return s
}
