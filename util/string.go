package util

import (
	"fmt"
)

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
		fmt.Printf("[%q][%q]\n", c[0], c[len(c)-1])
		if (c[0] == '"' && c[len(c)-1] == '"') || (c[0] == '\'' && c[len(c)-1] == '\'') {
			fmt.Printf("[%s]\n", c[1:len(c)-1])
			s[i] = c[1:len(c)-1]
		}
	}

	return s
}
