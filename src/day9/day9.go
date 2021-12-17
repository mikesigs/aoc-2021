package day9

import (
	"fmt"
	"mikesigs/aoc-2021/src/shared"
)

type IPoint interface {
	Basin() int
	SetBasin(int)
	String()
}

type Point struct {
	Value                     int
	Left, Right, Upper, Lower *Point
	basin                     int
}

func NewPoint(value int) *Point {
	return &Point{Value: value, basin: -1}
}

func (p *Point) String() string {
	return fmt.Sprintf("Value: %c, Basin: %d\n", p.Value, p.basin)
}

func (p *Point) Basin() int {
	return p.basin
}

func (p *Point) SetBasin(b int) {
	p.basin = b
}

func (p *Point) IsLowerThanNeighbours() bool {
	return (p.Left == nil || p.Value < p.Left.Value) &&
		(p.Right == nil || p.Value < p.Right.Value) &&
		(p.Upper == nil || p.Value < p.Upper.Value) &&
		(p.Lower == nil || p.Value < p.Lower.Value)
}

func Part1() int {
	lines, err := shared.ReadLines("day9.txt")
	shared.Check(err)
	result := 0

	grid := newGrid(lines)

	for y := range grid {
		for _, p := range grid[y] {
			// fmt.Printf("(%d,%d): p:%c,l:%c,r:%c,u:%c,d:%c\n", y, x, p, l, r, u, d)
			if p.IsLowerThanNeighbours() {
				result += p.Value + 1
			}
		}
	}
	return result
}

// func Part2() int {
// 	lines, err := shared.ReadLines("day9.txt")
// 	shared.Check(err)

// 	grid := newGrid(lines)
// 	for y := range grid {
// 		for x := range grid[y] {
// 			var p, l, u byte
// 			p = lines[y][x]
// 			if x > 0 {
// 				l = lines[y][x-1]
// 			}
// 			if y > 0 {
// 				u = lines[y-1][x]
// 			}

// 			if l != 0 && l != '9' {
// 				p.SetBasin(l.Basin())
// 			}

// 		}
// 	}
// 	return 0
// }

func newGrid(lines []string) [][]Point {
	grid := make([][]Point, len(lines))
	for y := range lines {
		grid[y] = make([]Point, len(lines[y]))
		for x := range lines[y] {
			char := lines[y][x]
			grid[y][x] = *NewPoint(btoi(char))
			p := &grid[y][x]
			if x > 0 {
				l := &grid[y][x-1]
				p.Left = l
				l.Right = p
			}
			if y > 0 {
				u := &grid[y-1][x]
				p.Upper = u
				u.Lower = p
			}
		}
	}
	return grid
}

func btoi(b byte) int {
	return int(b - '0')
}
