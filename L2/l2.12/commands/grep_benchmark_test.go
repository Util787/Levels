package commands

import (
	"strings"
	"testing"
)

const inputSize = 1000000

func generateInput(n int) []string {
	input := make([]string, n)
	for i := 0; i < n; i++ {
		if i%10 == 0 {
			input[i] = "match this line"
		} else {
			input[i] = "some other text"
		}
	}
	return input
}

func BenchmarkSequentialFindMatchIndexes(b *testing.B) {
	lines := generateInput(inputSize)
	matchFunc := func(line string) bool {
		return strings.Contains(line, "match")
	}
	for i := 0; i < b.N; i++ {
		seqFindMatchIndexes(lines, matchFunc)
	}
}

func BenchmarkConcurrentFindMatchIndexes(b *testing.B) {
	lines := generateInput(inputSize)
	matchFunc := func(line string) bool {
		return strings.Contains(line, "match")
	}
	for i := 0; i < b.N; i++ {
		findMatchIndexes(lines, matchFunc)
	}
}

func seqFindMatchIndexes(lines []string, matchFunc func(string) bool) []int {
	matchIndexes := make([]int, 0, 25)
	for i, line := range lines {
		match := matchFunc(line)
		if match {
			matchIndexes = append(matchIndexes, i)
		}
	}
	return matchIndexes
}
