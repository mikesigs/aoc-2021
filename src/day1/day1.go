package day1

import (
	"mikesigs/aoc-2021/src/shared"
)

func Part1() int {
	nums := shared.ReadNums("/workspace/aoc-2021/data/day1.txt")

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
	nums := shared.ReadNums("/workspace/aoc-2021/data/day1.txt")

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

func sum(nums []int) int {
	result := 0
	for _, n := range nums {
		result += n
	}
	return result
}
