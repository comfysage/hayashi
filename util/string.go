package util

import "strings"

func StringSplit(str string) []string {
	splits := make([]int, 0)
	quoted := false
	var quotec rune
	for i, s := range str {
		if quoted {
			if s == quotec {
				quoted = false
				continue
			}
			continue
		}

		if s == '"' || s == '\'' {
			quoted = true
			quotec = s
			continue
		}
		if s == ' ' {
			splits = append(splits, i)
		}
	}
	splits = append(splits, len(str))

	i := 0
	s := []string{}
	for _, j := range splits {
		sl := str[i:j]
		if sl != "" {
			s = append(s, sl)
		}
		i = j+1
	}
	for i, c := range s {
		if len(c) < 2 {
			continue
		}
		if (c[0] == '"' && c[len(c)-1] == '"') || (c[0] == '\'' && c[len(c)-1] == '\'') {
			s[i] = c[1:len(c)-1]
		}
	}

	return s
}

func SearchStrings(arr []string, input string) []uint {
	var matches []uint

	for i, str := range arr {
		if strings.Contains(str, input) {
			matches = append(matches, uint(i))
		}
	}

	return matches
}
