package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/pflag"
)

var sortUsage string = "\nUsage: go run sort.go [-k key] [-n] [-r] [-u] <input file> <output file>\n"

func main() {
	SortCmd()
}

type sortFlags struct {
	kFlag *int
	nFlag *bool
	rFlag *bool
	uFlag *bool
}

func getSortArgs() (flags *sortFlags, inputPath string, outputPath string) {
	flags = &sortFlags{
		kFlag: pflag.IntP("key", "k", 0, "Sort by column number (starting from 1)"),
		nFlag: pflag.BoolP("numeric", "n", false, "Sort by numerical value"),
		rFlag: pflag.BoolP("reverse", "r", false, "Sort in reverse order"),
		uFlag: pflag.BoolP("unique", "u", false, "Output only unique lines"),
	}

	pflag.Parse()
	if *flags.kFlag < 0 {
		fmt.Println("Warning: -k (key) flag cannot be negative. Setting to 1.")
		*flags.kFlag = 1
	}
	fmt.Println("Flags:")
	fmt.Printf(" -k (key): %d\n", *flags.kFlag)
	fmt.Printf(" -n (numeric): %v\n", *flags.nFlag)
	fmt.Printf(" -r (reverse): %v\n", *flags.rFlag)
	fmt.Printf(" -u (unique): %v\n", *flags.uFlag)

	inputPath = os.Args[len(os.Args)-2]
	fmt.Println("Input file:", inputPath)

	outputPath = os.Args[len(os.Args)-1]
	fmt.Println("Output file:", outputPath)
	fmt.Println()

	return flags, inputPath, outputPath
}

// SortCmd is command that sorts lines from input file and writes result to output file. It gets arguments from terminal.
func SortCmd() {
	flags, inputPath, outputPath := getSortArgs()

	file, err := os.Open(inputPath)
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("error opening file: %s\n", err.Error()))
		fmt.Println(sortUsage)
		os.Exit(1)
	}
	defer file.Close()

	lines := make([]string, 0, 250)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)

	}
	if err := scanner.Err(); err != nil {
		os.Stderr.WriteString(fmt.Sprintf("error reading file: %s\n", err.Error()))
		os.Exit(1)
	}

	sortedLines, errs := sortLines(flags, lines)
	if errs != nil {
		os.Stderr.WriteString("errors during sorting:\n")
		for _, e := range errs {
			os.Stderr.WriteString(fmt.Sprintf("\t %s\n", e.Error()))
		}
		os.Exit(1)
	}

	err = writeOutput(sortedLines, outputPath)
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("%s\n", err.Error()))
		os.Exit(1)
	}
}

func sortLines(flags *sortFlags, lines []string) ([]string, []error) {
	if *flags.uFlag {
		lines = makeStrSliceUnique(lines) // call before sort. If called after sort it will break the sorted order
	}

	var sortErrs []error
	sort.Slice(lines, func(i, j int) bool {
		a, b := lines[i], lines[j] // elements to sort are a and b for convenience

		if *flags.kFlag > 0 {
			colsA := strings.Fields(a)
			colsB := strings.Fields(b)

			// validation
			if *flags.kFlag-1 < len(colsA) && *flags.kFlag-1 < len(colsB) {
				a = colsA[*flags.kFlag-1]
				b = colsB[*flags.kFlag-1]
			} else {
				// if beyond limit, sort by last column
				a = colsA[len(colsA)-1]
				b = colsB[len(colsB)-1]
			}
		}

		if *flags.nFlag {
			afloat, errA := strconv.ParseFloat(a, 64)
			if errA != nil {
				sortErrs = append(sortErrs, fmt.Errorf("error converting string to float: %w", errA))
			}
			bfloat, errB := strconv.ParseFloat(b, 64)
			if errB != nil {
				sortErrs = append(sortErrs, fmt.Errorf("error converting string to float: %w", errB))
			}

			if errA == nil && errB == nil {
				return afloat < bfloat
			}
		}

		return strings.ToLower(a) < strings.ToLower(b)
	})
	if sortErrs != nil {
		return nil, sortErrs
	}

	if *flags.rFlag {
		reverseStrSlice(lines)
	}

	return lines, nil
}

func reverseStrSlice(s []string) {
	i := 0
	j := len(s) - 1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

func makeStrSliceUnique(sl []string) []string {
	unique := make([]string, 0, len(sl))
	m := make(map[string]struct{})

	for _, str := range sl {
		if _, ok := m[str]; !ok {
			m[str] = struct{}{}
			unique = append(unique, str)
		}
	}

	return unique
}

func writeOutput(lines []string, outputPath string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error creating output file: %w", err)
	}
	defer file.Close()

	for _, line := range lines {
		if _, err := file.WriteString(line + "\n"); err != nil {
			return fmt.Errorf("error writing to output file: %w", err)
		}
	}

	return nil
}
