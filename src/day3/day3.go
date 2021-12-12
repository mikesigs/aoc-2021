package day3

import (
	"mikesigs/aoc-2021/src/shared"
	"strconv"
)

func Part1() int64 {
	lines, err := shared.ReadLines("day3.txt")
	shared.Check(err)

	sums := make([]int, len(lines[0]))
	for _, line := range lines {
		for i, r := range line {
			if r == '1' {
				sums[i]++
			}
		}
	}

	var gammaString, epsilonString string
	for _, x := range sums {
		if x >= len(lines)/2 {
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

func Part2() int64 {
	lines, err := shared.ReadLines("day3.txt")
	shared.Check(err)

	o2Rating := filter(lines, func(s []string, i int) byte { mostCommon, _ := getBitCommonalityAtPosition(s, i); return mostCommon })
	co2Rating := filter(lines, func(s []string, i int) byte { _, leastCommon := getBitCommonalityAtPosition(s, i); return leastCommon })

	lifeSupportRating := o2Rating * co2Rating
	return lifeSupportRating
}

func filter(readings []string, getBitCriterion func([]string, int) byte) int64 {
	pos := 0
	for pos < len(readings[0]) {
		criterion := getBitCriterion(readings, pos)
		// fmt.Printf("pos %d criterion %s\n", pos, string(criterion))

		readings = filterByBitCriterion(readings, pos, criterion)
		// fmt.Printf("readings: %v\n", readings)

		if len(readings) == 1 {
			rating, err := strconv.ParseInt(readings[0], 2, 26)
			shared.Check(err)
			return rating
		} else if len(readings) == 0 {
			panic("Unexpectedly filtered all values")
		}

		pos++
	}

	panic("Failed to filter list to unique value")
}

func filterByBitCriterion(lines []string, pos int, criterion byte) []string {
	var keep []string
	for _, line := range lines {
		if line[pos] == criterion {
			keep = append(keep, line)
		}
	}
	return keep
}

func getBitCommonalityAtPosition(lines []string, pos int) (byte, byte) {
	numOnes := 0.0
	for _, line := range lines {
		if line[pos] == '1' {
			numOnes++
		}
	}

	if numOnes >= float64(len(lines))/2.0 {
		return '1', '0'
	} else {
		return '0', '1'
	}
}
