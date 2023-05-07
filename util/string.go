package util

func StringSplit(str string) []string {
	splits := make([]int, 0)
	quoted := false
	var quotec rune
	for i, s := range str {
		if quoted {
			if s == quotec {
				splits = append(splits, i)
				quoted = false
				continue
			}
			continue
		}

		if s == '"' || s == '\'' {
			quoted = true
			quotec = s
			splits = append(splits, i)
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
		i = j + 1
	}

	return s
}
