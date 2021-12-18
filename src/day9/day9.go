package day9

import (
	"mikesigs/aoc-2021/src/day9/grid"
	"mikesigs/aoc-2021/src/shared"
)

func Part1() int {
	lines, err := shared.ReadLines("day9.txt")
	shared.Check(err)

	grid := grid.NewGrid(lines)

	return grid.SumOfRiskLevels()
}

func Part2() int {
	lines, err := shared.ReadLines("day9.txt")
	shared.Check(err)

	grid := grid.NewGrid(lines)
	grid.DefineBasins()
	// grid.PrintBasinMap()

	return grid.ProductOfThreeLargestBasins()
}
