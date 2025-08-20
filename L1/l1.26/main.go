package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(RepeatsCheck("abcd"))
	fmt.Println(RepeatsCheck("abCdefAaf"))
	fmt.Println(RepeatsCheck("aabcd"))
}

func RepeatsCheck(s string) bool {
	s = strings.ToLower(s)

	m := make(map[rune]uint8)

	for _, val := range s { // автоматически первращается в слайс рун
		m[val]++
	}

	for _, val := range m {
		if val > 1 {
			return false
		}
	}

	return true
}
