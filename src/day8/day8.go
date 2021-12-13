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
)

func Part2() int {
	lines, err := shared.ReadLines("day8.txt")
	shared.Check(err)

	for _, line := range lines {
		digits := map[int]segment7.Pattern{}
		// Get raw patterns from line input and sort them by length
		rawPatterns := strings.Fields(strings.Split(line, " | ")[0])
		sort.SliceStable(rawPatterns, func(i, j int) bool { return len(rawPatterns[i]) < len(rawPatterns[j]) })

		// Create a list of the patterns
		patterns := make([]segment7.Pattern, numPatterns)
		for i, p := range rawPatterns {
			patterns[i] = segment7.NewPattern(p)
		}

		fmt.Println(patterns)

		// digits 1,7,4 are in patterns 0,1,2 respectively, once sorted by length
		digits[1] = patterns[0] // cf
		digits[4] = patterns[2] // bcdf
		digits[7] = patterns[1] // acf

		// We can now determine the 6-signal digits with some reductive logic
		// The 6-signal digits are in patterns 6,7,8, which must be the digits 0,6,9
		for _, p := range patterns[6:9] {
			if p.NonMatchingSignalCount(digits[1].Signals) == 5 {
				digits[6] = p
			} else if p.NonMatchingSignalCount(digits[4].Signals) == 2 {
				digits[9] = p
			} else {
				digits[0] = p
			}
		}

		fmt.Println(digits)

		// 1: be 
		// 7: bde 
		// 4: bceg 
		// cdefg 
		// bcdef 
		// abcdf 
				
		// bcdefg-bceg=df=2 !
		// acdefg-bceg=adf=3 
		// abdefg-bceg=adf=3
		
		// bcdefg-be=cdfg=4
		// acdefg-be=acdfg=5 ! 
		// abdefg-be=adfg=4 
		
		// bcdefg-bde=cfg=3
		// acdefg-bde=acfg=4 !
		// abdefg-bde=afg=3 
		

		// abcdefg

		// We can now determine digits 0 and 6 by removing known signals acf+g
		// 0: abcefg-acf=beg
		// 6: abdefg-acf=bdeg, if len(signal.Except(digit_7)) == 4 { signal.SetDigit(6) }
		// 9: abcdfg-acf=bdg

		// 0: abcefg-cf=abeg
		// 6: abdefg-cf=abdeg, if len(signal.Except(digit_1)) == 5 { signal.SetDigit(6) }
		// 9: abcdfg-cf=abdg

		// 0: abcefg-bcdf=aeg
		// 6: abdefg-bcdf=aeg
		// 9: abcdfg-bcdf=ag, if len(signal.Except(digit_4)) == 2 { signal.SetDigit(9) }

		// // pattern[3,4,5] are always length 5 so must be one of Digits 2,3,5
		// // sixDigitPatterns := patterns[6:9] // 0, 6, or 9
		// // for _, p := range sixDigitPatterns {
		// // 	digit := signal.Signal{Value: p}
		// // 2: acdeg-abcdf = eg
		// // 3: acdfg-a  = cdfg, cdfg-cf = dg, dg-bd = g
		// // 5: abdfg-a=bdfg,
		// // }

		// // The last pattern is always Digit 8

		// outputs := strings.Fields(strings.Split(line, " | ")[1])
	}

	return 0
}
