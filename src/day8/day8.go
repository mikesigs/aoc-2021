package day8

import (
	"fmt"
	"mikesigs/aoc-2021/src/shared"
	"strings"
)

var digits = map[string]int{
	"abcefg":  0,
	"cf":      1,
	"acdeg":   2,
	"acdfg":   3,
	"bcdf":    4,
	"abdfg":   5,
	"abdefg":  6,
	"acf":     7,
	"abcdefg": 8,
	"abcdfg":  9,
}

func Part1() int {
	lines, err := shared.ReadLines("day8.txt")
	shared.Check(err)
	readings := strings.Fields(lines[0])
	for _, r := range readings {
		sortedReading := sortString(r)
		matchingDigit := digits[sortedReading]
		fmt.Printf("reading: %s, sorted: %s, match: %d\n", r, sortedReading, matchingDigit)
	}

	return 0
}

func sortString(input string) string {
	var output string
	for r := 'a'; r <= 'g'; r++ {
		if strings.ContainsRune(input, r) {
			output += string(r)
		}
	}
	return output
}
