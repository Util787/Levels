package commands

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"strings"
	"sync"

	"github.com/spf13/pflag"
)

const grepUsage string = "\nUsage: go run g.go [options] <pattern> [<input file>] <output file>"

const numParts = 4 // number of parts to split the lines into for concurrent search

type grepFlags struct {
	AFlag *int
	BFlag *int
	CFlag *int
	cFlag *bool
	iFlag *bool
	vFlag *bool
	FFlag *bool
	nFlag *bool
}

func getGrepArgs() (flags *grepFlags, pattern string, inputPath string, outputPath string) {
	flags = &grepFlags{
		AFlag: pflag.IntP("after-context", "A", 0, "Print N lines after each match"),
		BFlag: pflag.IntP("before-context", "B", 0, "Print N lines before each match"),
		CFlag: pflag.IntP("context", "C", 0, "Print N lines before and after each match"),
		cFlag: pflag.BoolP("count", "c", false, "Print only the count of matching lines"),
		iFlag: pflag.BoolP("ignore-case", "i", false, "Ignore letter case"),
		vFlag: pflag.BoolP("invert-match", "v", false, "Select non-matching lines"),
		FFlag: pflag.BoolP("fixed-strings", "F", false, "Interpret pattern as a fixed string"),
		nFlag: pflag.BoolP("line-number", "n", false, "Print line number with output lines"),
	}

	pflag.Parse()
	args := pflag.Args()

	if len(args) == 1 {
		pattern = args[0]
		return
	}
	if len(args) == 2 {
		pattern = args[0]
		inputPath = args[1]
		return
	}
	if len(args) == 3 {
		pattern = args[0]
		inputPath = args[1]
		outputPath = args[2]
		return
	}

	fmt.Fprintf(os.Stderr, "error: Invalid arguments")
	fmt.Println(grepUsage)
	os.Exit(1)
	return
}

// GrepCmd is command that searches for pattern in input file and writes result to output file(my implementation of unix grep). It gets arguments from terminal.
func GrepCmd() {
	flags, pattern, inputPath, outputPath := getGrepArgs()
	matchFunc := buildMatchFunc(flags, pattern)

	lines, output := getLinesAndOutput(inputPath, outputPath)
	defer output.Close()

	before := *flags.BFlag
	after := *flags.AFlag
	if *flags.CFlag > 0 {
		before = *flags.CFlag
		after = *flags.CFlag
	}

	matchIndexes := findMatchIndexes(lines, matchFunc)

	if *flags.cFlag {
		fmt.Fprintln(output, len(matchIndexes))
		return
	}

	// Context line handling
	printed := make(map[int]struct{})
	for _, idx := range matchIndexes {
		start := idx - before
		if start < 0 {
			start = 0
		}
		end := idx + after
		if end >= len(lines) {
			end = len(lines) - 1
		}

		// writing output
		for i := start; i <= end; i++ {
			if _, ok := printed[i]; ok {
				continue
			}
			printed[i] = struct{}{}
			if *flags.nFlag {
				fmt.Fprintf(output, "%d: %s\n", i+1, lines[i])
			} else {
				fmt.Fprintln(output, lines[i])
			}
		}
	}
}

func buildMatchFunc(flags *grepFlags, pattern string) func(line string) bool {
	if *flags.FFlag {
		// fixed string match
		if *flags.iFlag {
			return func(line string) bool {
				if *flags.vFlag {
					return !fixedPatternMatch(strings.ToLower(line), strings.ToLower(pattern))
				}
				return fixedPatternMatch(strings.ToLower(line), strings.ToLower(pattern))
			}
		}
		return func(line string) bool {
			if *flags.vFlag {
				return !fixedPatternMatch(line, pattern)
			}
			return fixedPatternMatch(line, pattern)
		}
	}

	// regex match
	var re *regexp.Regexp
	var err error
	if *flags.iFlag {
		re, err = regexp.Compile("(?i)" + pattern)
	} else {
		re, err = regexp.Compile(pattern)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: Invalid regex pattern: %v\n", err)
		os.Exit(1)
	}
	return func(line string) bool {
		if *flags.vFlag {
			return !re.MatchString(line)
		}
		return re.MatchString(line)
	}

}

func fixedPatternMatch(line, pattern string) bool {
	seq := strings.FieldsSeq(line)
	for word := range seq {
		if word == pattern {
			return true
		}
	}
	return false
}

func getLinesAndOutput(inputPath, outputPath string) ([]string, *os.File) {
	var input *os.File
	var output *os.File
	var err error

	if inputPath == "" {
		input = os.Stdin
	} else {
		input, err = os.Open(inputPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error opening input file: %v\n", err)
			os.Exit(1)
		}
		defer input.Close()
	}

	lines := make([]string, 0, 250)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		os.Stderr.WriteString(fmt.Sprintf("error reading file: %s\n", err.Error()))
		os.Exit(1)
	}

	if outputPath == "" {
		output = os.Stdout
	} else {
		output, err = os.Create(outputPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error creating output file: %v\n", err)
			os.Exit(1)
		}
	}

	return lines, output
}

// the idea is to split the lines into parts and search each part concurrently, then find match lines with their global index, then merge the results and sort it.
func findMatchIndexes(lines []string, matchFunc func(string) bool) []int {
	parts, interval := splitLines(lines, numParts)

	matchCh := make(chan []int, numParts)
	wg := &sync.WaitGroup{}

	for partIdx, part := range parts {
		wg.Add(1)
		go func() {
			defer wg.Done()
			localMatches := make([]int, 0, 10)
			for lineIdx, line := range part {
				if matchFunc(line) {
					localMatches = append(localMatches, lineIdx+partIdx*interval) // to get global line index
				}
			}
			matchCh <- localMatches
		}()
	}
	wg.Wait()
	close(matchCh)

	matchIndexes := make([]int, 0, 25)
	for localMatches := range matchCh {
		matchIndexes = append(matchIndexes, localMatches...)
	}
	slices.Sort(matchIndexes)

	return matchIndexes
}

func splitLines(lines []string, numParts int) ([][]string, int) {
	interval := int(math.Ceil(float64(len(lines)) / float64(numParts)))
	split := make([][]string, 0, numParts)

	for i := 0; i < len(lines); i += interval {
		end := i + interval
		if end > len(lines) {
			end = len(lines)
		}
		split = append(split, lines[i:end])
	}

	return split, interval
}
