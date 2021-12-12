package day7

import (
	"math"
	"mikesigs/aoc-2021/src/shared"
)

func Part1() int {
	lines, err := shared.ReadLines("day7.txt")
	shared.Check(err)
	crabs := shared.SplitInts(lines[0], ",")
	crabPositions, maxPosition := getCrabPositions(crabs)
	return cheapestMovementCost(crabPositions, maxPosition, part1Move)
}

func Part2() int {
	lines, err := shared.ReadLines("day7.txt")
	shared.Check(err)
	crabs := shared.SplitInts(lines[0], ",")
	crabPositions, maxPosition := getCrabPositions(crabs)
	return cheapestMovementCost(crabPositions, maxPosition, part2Move)
}

func part1Move(start, dest int) int {
	return shared.Abs(start - dest)
}

func part2Move(pos, dest int) int {
	dist := shared.Abs(pos - dest)
	var cost int
	for steps := 1; steps <= dist; steps++ {
		cost += steps
	}

	return cost
}

func cheapestMovementCost(crabPositions map[int]int, maxPosition int, moveFunc func(int, int) int) int {
	var cheapest int = math.MaxInt
	for dest := 0; dest <= maxPosition; dest++ {
		var cost int
		for pos, numCrabs := range crabPositions {
			if numCrabs == 0 {
				continue
			}

			cost += numCrabs * moveFunc(pos, dest)
		}
		if cost < cheapest {
			cheapest = cost
		}
	}
	return cheapest
}

func getCrabPositions(crabs []int) (map[int]int, int) {
	crabPositions := map[int]int{}
	var maxPosition int
	for _, p := range crabs {
		crabPositions[p]++
		if p > maxPosition {
			maxPosition = p
		}
	}
	return crabPositions, maxPosition
}
