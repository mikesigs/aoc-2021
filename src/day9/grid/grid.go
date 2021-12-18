package grid

import (
	"fmt"
	"sort"
	"strconv"
)

const (
	UndefinedBasin int = -1
)

type IPoint interface {
	TraverseBasin(basin int)
	String()
}

type Point struct {
	Value                 int
	Left, Right, Up, Down *Point
	Basin                 int
}

func NewPoint(value int) *Point {
	return &Point{Value: value, Basin: UndefinedBasin}
}

func (p *Point) TraverseBasin(basin int) {
	var traverse func(*Point, int)

	traverse = func(p *Point, basin int) {
		setBasin := func(p *Point, basin int) {
			if p != nil && p.Value != 9 && p.Basin == UndefinedBasin {
				p.Basin = basin
				traverse(p, basin)
			}
		}
		setBasin(p.Left, basin)
		setBasin(p.Right, basin)
		setBasin(p.Up, basin)
		setBasin(p.Down, basin)
	}

	traverse(p, basin)
}

func (p *Point) String() string {
	s := "Map:\n"
	var l, r, lpad, rpad string
	if p.Left != nil {
		lpad = " "
		l = strconv.Itoa(p.Left.Value)
	}
	if p.Right != nil {
		rpad = " "
		r = strconv.Itoa(p.Right.Value)
	}

	if p.Up != nil {
		s += fmt.Sprintf("%s%d%s\n", lpad, p.Up.Value, rpad)
	}
	s += fmt.Sprintf("%s%d%s\n", l, p.Value, r)
	if p.Down != nil {
		s += fmt.Sprintf("%s%d%s\n", lpad, p.Down.Value, rpad)
	}
	s += fmt.Sprintf("Basin: %d\n", p.Basin)
	return s
}

func (p *Point) IsLowerThanNeighbours() bool {
	return (p.Left == nil || p.Value < p.Left.Value) &&
		(p.Right == nil || p.Value < p.Right.Value) &&
		(p.Up == nil || p.Value < p.Up.Value) &&
		(p.Down == nil || p.Value < p.Down.Value)
}

type IGrid interface {
	SetBasins()
	PrintBasinMap()
}

type Grid struct {
	grid [][]Point
}

func (g *Grid) SumOfRiskLevels() int {
	result := 0
	for y := range g.grid {
		for _, p := range g.grid[y] {
			if p.IsLowerThanNeighbours() {
				result += p.Value + 1
			}
		}
	}
	return result
}

func (g *Grid) ProductOfThreeLargestBasins() int {
	basinSizes := map[int]int{}
	for y := range g.grid {
		for _, p := range g.grid[y] {
			if p.Basin != UndefinedBasin {
				basinSizes[p.Basin]++
			}
		}
	}

	sizes := make([]int, len(basinSizes))
	i := 0
	for _, s := range basinSizes {
		sizes[i] = s
		i++
	}
	sort.Ints(sizes)
	product := 1
	for _, s := range sizes[len(sizes)-3:] {
		product *= s
	}
	return product
}

func (g *Grid) DefineBasins() {
	basin := 1
	for y := range g.grid {
		for x := range g.grid[y] {
			p := &g.grid[y][x]
			if p.Value != 9 && p.Basin == UndefinedBasin {
				p.TraverseBasin(basin)
				basin++
			}
		}
	}
}

func (g *Grid) PrintBasinMap() {
	for y := range g.grid {
		for _, p := range g.grid[y] {
			if p.Basin > 0 {
				fmt.Printf("%03d", p.Basin)
			} else {
				fmt.Printf("   ")
			}
		}
		fmt.Println()
	}
}

func NewGrid(lines []string) *Grid {
	grid := make([][]Point, len(lines))
	for y := range lines {
		grid[y] = make([]Point, len(lines[y]))
		for x := range lines[y] {
			char := lines[y][x]
			grid[y][x] = *NewPoint(int(char - '0'))
			p := &grid[y][x]
			if x > 0 {
				l := &grid[y][x-1]
				p.Left = l
				l.Right = p
			}
			if y > 0 {
				u := &grid[y-1][x]
				p.Up = u
				u.Down = p
			}
		}
	}
	return &Grid{grid}
}
