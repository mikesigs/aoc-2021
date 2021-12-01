package day1

import (
	"bufio"
	"os"
	"strconv"
)

func Part1() int {
	nums := readNums("/workspace/aoc-2021/data/day1.txt")

	var prev int
	higher := -1
	for _, curr := range nums {
		if higher == -1 {
			higher = 0
		} else if curr > prev {
			higher += 1
		}

		prev = curr
	}

	return higher
}

func Part2() int {
	nums := readNums("/workspace/aoc-2021/data/day1.txt")

	var prev int
	higher := -1
	for i := range nums {
		upper := i + 3
		if upper > len(nums) {
			upper = len(nums)
		}
		s := nums[i:upper]

		curr := sum(s)

		if higher == -1 {
			higher = 0
		} else if curr > prev {
			higher += 1
		}

		prev = curr
	}

	return higher
}

func sum(nums []int) int {
	result := 0
	for _, v := range nums {
		result += v
	}
	return result
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func readNums(path string) []int {
	lines, err := readLines(path)
	check(err)
	var nums []int
	for _, line := range lines {
		num, err := strconv.Atoi(line)
		check(err)
		nums = append(nums, num)
	}
	return nums
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
