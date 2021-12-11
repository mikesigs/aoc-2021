package day6

import (
	"mikesigs/aoc-2021/src/shared"
	"strconv"
	"strings"
)

func Part1() int {
	lines, err := shared.ReadLines("/workspace/aoc-2021/data/day6.txt")
	shared.Check(err)
	fish := loadFish(lines[0])

	for i := 0; i < 40; i++ {
		fish = generationA(fish)
	}

	return len(fish)
}

func Part2() int {
	lines, err := shared.ReadLines("/workspace/aoc-2021/data/day6.txt")
	shared.Check(err)

	fish := make(map[int]int, 9)
	for _, d := range loadFish(lines[0]) {
		fish[d]++
	}

	for d := 0; d < 256; d++ {
		today := d % 9
		fish[(today+7)%9] += fish[today]
	}

	var sum int
	for i := range fish {
		sum += fish[i]
	}

	return sum
}

func loadFish(input string) []int {
	inputs := strings.Split(input, ",")
	nums := make([]int, len(inputs))
	for i, x := range inputs {
		n, err := strconv.Atoi(x)
		shared.Check(err)
		nums[i] = n
	}

	return nums
}

func generationA(fish []int) []int {
	for i := range fish {
		if fish[i] == 0 {
			fish[i] = 6
			fish = append(fish, 8)
		} else {
			fish[i]--
		}
	}
	return fish
}
