package hw02unpackstring

import (
	"errors"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	runesList := []rune(str)
	var builder strings.Builder
	i := 0
	for i < len(runesList) {
		if '0' <= runesList[i] && runesList[i] <= '9' {
			return "", ErrInvalidString
		} else if runesList[i] == '\\' {
			i++
		}

		window := runesList[i:min(i+2, len(runesList))]

		switch {
		case len(window) == 1:
			builder.WriteRune(window[0])
		case '0' <= window[1] && window[1] <= '9':
			count := int(window[1] - '0')
			builder.WriteString(strings.Repeat(string(window[0]), count))
			i++
		default:
			builder.WriteRune(window[0])
		}

		i++
	}
	return builder.String(), nil
}
