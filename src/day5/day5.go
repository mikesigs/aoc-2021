package day5

import (
	"fmt"
	"mikesigs/aoc-2021/src/shared"
)

type coord struct {
	x, y int
}

func Part1() int {
	lines, err := shared.ReadLines("day5.txt")
	shared.Check(err)

	grid := make(map[coord]int)

	for _, line := range lines {
		var x1, y1, x2, y2 int
		fmt.Sscanf(line, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		// fmt.Printf("%3d,%3d -> %3d,%3d\n", x1, y1, x2, y2)
		if x1 == x2 {
			x := x1
			if y1 <= y2 {
				for y := y1; y <= y2; y++ {
					grid[coord{x, y}]++
				}
			} else {
				for y := y1; y >= y2; y-- {
					grid[coord{x, y}]++
				}
			}
		} else if y1 == y2 {
			y := y1
			if x1 <= x2 {
				for x := x1; x <= x2; x++ {
					grid[coord{x, y}]++
				}
			} else {
				for x := x1; x >= x2; x-- {
					grid[coord{x, y}]++
				}
			}
		}
	}
	// fmt.Println(grid)

	// for y := 0; y < 9; y++ {
	// 	for x := 0; x < 9; x++ {
	// 		value := grid[coord{x, y}]
	// 		if value == 0 {
	// 			fmt.Print(".")
	// 		} else {
	// 			fmt.Printf("%d", value)
	// 		}
	// 	}
	// 	fmt.Println("")
	// }

	var count int
	for _, elem := range grid {
		if elem >= 2 {
			count++
		}
	}

	return count
}

func Part2() int {
	lines, err := shared.ReadLines("day5.txt")
	shared.Check(err)

	grid := make(map[coord]int)

	for _, line := range lines {
		var x1, y1, x2, y2 int
		fmt.Sscanf(line, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		// fmt.Printf("%d,%d -> %d,%d\n", x1, y1, x2, y2)
		rise := y2 - y1
		run := x2 - x1
		x := x1
		y := y1
		for {
			grid[coord{x, y}]++
			if x == x2 && y == y2 {
				break
			}

			if rise > 0 {
				y++
			} else if rise < 0 {
				y--
			}
			if run > 0 {
				x++
			} else if run < 0 {
				x--
			}
		}
	}
	// fmt.Println(grid)

	// for y := 0; y <= 9; y++ {
	// 	for x := 0; x <= 9; x++ {
	// 		value := grid[coord{x, y}]
	// 		if value == 0 {
	// 			fmt.Print(".")
	// 		} else {
	// 			fmt.Printf("%d", value)
	// 		}
	// 	}
	// 	fmt.Println("")
	// }

	var count int
	for _, elem := range grid {
		if elem >= 2 {
			count++
		}
	}

	return count
}
