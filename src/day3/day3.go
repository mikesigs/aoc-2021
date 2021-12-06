package day3

import (
	"mikesigs/aoc-2021/src/shared"
	"strconv"
)

func Part1() int64 {
	lines, err := shared.ReadLines("/workspace/aoc-2021/data/day3.txt")
	shared.Check(err)

	count := 0
	sums := make([]int, len(lines[0]))
	for _, line := range lines {
		for i, r := range line {
			if r == '1' {
				sums[i]++
			}
		}
		count++
	}

	var gammaString, epsilonString string
	for _, x := range sums {
		if x >= count/2 {
			gammaString += "1"
			epsilonString += "0"
		} else {
			gammaString += "0"
			epsilonString += "1"
		}
	}
	gammaRate, err := strconv.ParseInt(gammaString, 2, 64)
	shared.Check(err)
	epsilonRate, err := strconv.ParseInt(epsilonString, 2, 64)
	shared.Check(err)

	return gammaRate * epsilonRate
}
