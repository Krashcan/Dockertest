package util

import(
	"unicode/utf8"
	"unicode"
)

func Reverse(input string) string{
	if input == ""{
		return ""
	}

	s := []rune(input)
	output := make([]rune, len(s))

	start := len(output)

	for i := 0; i < len(s); {
		if s[i] == utf8.RuneError{
			i++
			continue
		}
		j := i + 1

		for j < len(s) && ((unicode.Is(unicode.Mn, s[j])) || (unicode.Is(unicode.Me, s[j])) || (unicode.Is(unicode.Mc, s[j]))){
			j++
		}
		for k := j - 1; k >= i; k-- {
			start--
			output[start] = s[k]
		}

		i = j
	}

	return (string(output[start:]))
}