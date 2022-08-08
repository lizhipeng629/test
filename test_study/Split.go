package test_study

import "strings"

func Split(s, seq string) (result []string) {
	i := strings.Index(s, seq)
	for i > -1 {
		result = append(result, s[:i])
		s = s[i+len(seq):]
		i = strings.Index(s, seq)
	}
	result = append(result, s)
	return
}
