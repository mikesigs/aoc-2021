package day2

import (
	"fmt"
	"mikesigs/aoc-2021/src/shared"
)

func Part1() int {
	directions, err := shared.ReadLines("day2.txt")
	shared.Check(err)

	var dir string
	var amt int

	var x, y int
	for _, s := range directions {
		n, err := fmt.Sscanf(s, "%s %d", &dir, &amt)
		shared.Check(err)
		if n != 2 {
			panic("Unexpected number of args read by sscanf")
		}

		switch dir {
		case "forward":
			x += amt
		case "up":
			y -= amt
		case "down":
			y += amt
		default:
			panic("Unexpected direction")
		}
	}

	return x * y
}

func Part2() int {
	directions, err := shared.ReadLines("day2.txt")
	shared.Check(err)

	var dir string
	var x int

	var horizontal_pos, depth, aim int
	for _, s := range directions {
		n, err := fmt.Sscanf(s, "%s %d", &dir, &x)
		shared.Check(err)
		if n != 2 {
			panic("Unexpected number of args read by sscanf")
		}

		switch dir {
		case "down":
			aim += x
		case "up":
			aim -= x
		case "forward":
			horizontal_pos += x
			depth += aim * x
		default:
			panic("Unexpected direction")
		}
	}

	return horizontal_pos * depth
}
