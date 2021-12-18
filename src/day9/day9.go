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
	grid.PrintBasinMap()

	return grid.ProductOfThreeLargestBasins()
}


/*Actually, all the closures and functions I had there weren't necessary. After I refactored it all into a "class" it became
```func (p *Point) SetBasin(basin int) {
	if p != nil && p.Value != 9 && p.Basin == UndefinedBasin {
		p.Basin = basin
		p.Left.SetBasin(basin)
		p.Right.SetBasin(basin)
		p.Up.SetBasin(basin)
		p.Down.SetBasin(basin)
	}
}```*/