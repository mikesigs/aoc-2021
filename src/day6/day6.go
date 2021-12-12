package day6

import (
	"mikesigs/aoc-2021/src/shared"
)

func Part1() int {
	lines, err := shared.ReadLines("day6.txt")
	shared.Check(err)
	fish := shared.SplitInts(lines[0], ",")

	for i := 0; i < 40; i++ {
		fish = generationA(fish)
	}

	return len(fish)
}

func Part2() int {
	lines, err := shared.ReadLines("day6.txt")
	shared.Check(err)

	fish := make(map[int]int, 9)
	for _, d := range shared.SplitInts(lines[0], ",") {
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
