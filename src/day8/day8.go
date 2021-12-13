package day8

import (
	"fmt"
	"mikesigs/aoc-2021/src/day8/segment7"
	"mikesigs/aoc-2021/src/shared"
	"sort"
	"strings"
)

func Part1() int {
	lines, err := shared.ReadLines("day8.txt")
	shared.Check(err)
	var count int
	for _, line := range lines {
		outputValues := strings.Split(line, " | ")[1]
		outputs := strings.Fields(outputValues)
		for _, s := range outputs {
			if len(s) == 2 || len(s) == 3 || len(s) == 4 || len(s) == 7 {
				count++
			}
		}
	}

	return count
}

const (
	numPatterns = 10
	numOutcomes = 4
)

func Part2() int {
	lines, err := shared.ReadLines("day8.txt")
	shared.Check(err)

	var result int

	for _, line := range lines {
		parts := strings.Split(line, " | ")

		// Get raw patterns from line input and sort them by length
		rawPatterns := strings.Fields(parts[0])
		sort.SliceStable(rawPatterns, func(i, j int) bool { return len(rawPatterns[i]) < len(rawPatterns[j]) })

		// Map raw patterns to Pattern types
		patterns := make([]*segment7.Pattern, numPatterns)
		for i, p := range rawPatterns {
			patterns[i] = segment7.NewPattern(p)
		}

		patterns[0].Digit = 1
		patterns[1].Digit = 7
		patterns[2].Digit = 4
		patterns[9].Digit = 8

		digit_1 := patterns[0]
		digit_4 := patterns[2]

		// The 5-signal digits are in patterns 3,4,5, which must be the digits 2,3,5
		for _, p := range patterns[3:6] {
			if p.Subtract(digit_1).Length() == 3 {
				p.Digit = 3
			} else if p.Subtract(digit_4).Length() == 3 {
				p.Digit = 2
			} else {
				p.Digit = 5
			}
		}

		// The 6-signal digits are in patterns 6,7,8, which must be the digits 0,6,9
		for _, p := range patterns[6:9] {
			if p.Subtract(digit_1).Length() == 5 {
				p.Digit = 6
			} else if p.Subtract(digit_4).Length() == 2 {
				p.Digit = 9
			} else {
				p.Digit = 0
			}
		}

		// printPatterns("Patterns", patterns)

		// Get raw outcomes from line input and map to Pattern types
		rawOutcomes := strings.Fields(parts[1])
		outcomes := make([]*segment7.Pattern, numOutcomes)
		for i, o := range rawOutcomes {
			outcomes[i] = segment7.NewPattern(o)
		}

		// Match outcomes to deciphered digit patterns
		for _, o := range outcomes {
			for _, p := range patterns {
				if o.Signals == p.Signals {
					o.Digit = p.Digit
				}
			}
		}

		// printPatterns("Outcomes", outcomes)

		// Get 4-digit output value
		var value int
		for i := 0; i < len(outcomes); i++ {
			value += outcomes[i].Digit * pow(10, len(outcomes)-i-1)
		}
		// fmt.Println("Output", value)

		// Increment solution
		result += value
	}

	return result
}

func printPatterns(label string, patterns []*segment7.Pattern) {
	fmt.Printf("%s ", label)
	for _, p := range patterns {
		fmt.Printf("%d:%s ", p.Digit, p.Signals)
	}
	fmt.Println()
}

func pow(a, b int) int {
	result := 1
	for i := 0; i < b; i++ {
		result *= a
	}
	return result
}
