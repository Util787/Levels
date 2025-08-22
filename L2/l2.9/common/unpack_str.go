package common

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

// UnpackStr unpacks string where letters can be followed by digits.
// Each letter followed by a digit indicates how many times that letter should be repeated.
//
// Example:
//
// Input: "a2b3" Output: "aabbb"
func UnpackStr(s string) (string, error) {
	s = strings.TrimSpace(s)

	if len(s) == 0 {
		return "", errors.New("empty string")
	}

	runes := []rune(s)
	builder := strings.Builder{}

	if unicode.IsDigit(runes[0]) {
		return "", errors.New("first letter is digit")
	}

	for i := 0; i < len(runes); {

		if i+1 < len(runes) && unicode.IsDigit(runes[i+1]) {
			j := i + 1
			for j < len(runes) && unicode.IsDigit(runes[j]) { // чтобы обрабатывать если число >9
				j++ // j будет выходить за рамки 'только цифр' например при строке "f10b" он будет индексом "b" но дальше в слайсе отсечется
			}

			num, err := strconv.Atoi(string(runes[i+1 : j])) // собираю все цифры в число num
			if err != nil {
				return "", errors.New("failed to convert from string to number")
			}
			if num < 1 {
				return "", errors.New("cant use numbers less than 1")
			}

			for range num {
				builder.WriteRune(runes[i])
			}

			i = j // скипаю до следующей буквы
		} else {
			builder.WriteRune(runes[i])
			i++
		}
	}

	return builder.String(), nil
}
