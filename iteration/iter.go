package iteration

import "strings"

func Repeat(count int, a string) string {
	s := ""
	for i := 0; i<count; i++ {
		s += "a"
	}
	return s
}

func CheckString(c []string, s string) int {
	count := 0
	for i := 0; i<len(c); i++ {
		b := strings.Contains(s, c[i])
		if b {
			count++
		}
	}
	return count
}