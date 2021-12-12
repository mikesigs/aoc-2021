package day1

import (
	"mikesigs/aoc-2021/src/shared"
	"strconv"
)

func Part1() int {
	lines, err := shared.ReadLines("day1.txt")
	shared.Check(err)
	nums := readNums(lines)

	var prev int
	higher := -1
	for _, curr := range nums {
		if curr > prev {
			higher += 1
		}
		prev = curr
	}

	return higher
}

func Part2() int {
	lines, err := shared.ReadLines("day1.txt")
	shared.Check(err)
	nums := readNums(lines)

	var prev int
	higher := -1
	for i := 0; i <= len(nums)-3; i++ {
		curr := sum(nums[i : i+3])
		if curr > prev {
			higher += 1
		}
		prev = curr
	}
	return higher
}

func readNums(lines []string) []int {
	var nums []int
	for _, line := range lines {
		num, err := strconv.Atoi(line)
		shared.Check(err)
		nums = append(nums, num)
	}
	return nums
}

func sum(nums []int) int {
	result := 0
	for _, n := range nums {
		result += n
	}
	return result
}
