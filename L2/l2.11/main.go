package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	fmt.Println(groupAnagrams([]string{"пятак", "Пятка", "тяпка", "Листок", "слиток", "Столик", "стол"}))
}

func groupAnagrams(sl []string) map[string][]string {
	wordsGrMap := make(map[string][]string, len(sl))
	for _, word := range sl {
		word = strings.ToLower(word)

		runes := []rune(word)
		slices.Sort(runes)

		k := string(runes)
		wordsGrMap[k] = append(wordsGrMap[k], word)
	}

	res := make(map[string][]string, len(sl))
	for _, wordsGroup := range wordsGrMap {
		if len(wordsGroup) == 1 {
			continue
		}
		key := wordsGroup[0]    // первое встреченное слово
		slices.Sort(wordsGroup) // сортировка по возрастанию
		res[key] = wordsGroup
	}
	return res
}
