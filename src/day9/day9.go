package day9

import (
	"fmt"
	"mikesigs/aoc-2021/src/shared"
)

func Part1() int {
	lines, err := shared.ReadLines("day9.txt")
	shared.Check(err)

	// grid := make([][]int, len(lines))
	result := 0

	for y := range lines {
		// grid[y] = make([]int, len(lines[y]))
		for x := range lines[y] {
			var p, l, r, u, d byte
			p = lines[y][x]
			if x > 0 {
				l = lines[y][x-1]
			}
			if x < len(lines[y])-1 {
				r = lines[y][x+1]
			}
			if y > 0 {
				u = lines[y-1][x]
			}
			if y < len(lines)-1 {
				d = lines[y+1][x]
			}
			// fmt.Printf("(%d,%d): p:%c,l:%c,r:%c,u:%c,d:%c\n", y, x, p, l, r, u, d)
			if (l == 0 || p < l) &&
				(r == 0 || p < r) &&
				(u == 0 || p < u) &&
				(d == 0 || p < d) {
				v := int(p - '0')
				fmt.Printf("Found low point %d at {%d,%d}\n", v, x, y)
				fmt.Printf(" %c \n", u)
				fmt.Printf("%c%c%c\n", l, p, r)
				fmt.Printf(" %c\n", d)
				result += int(p-'0') + 1
			}
		}
	}
	return result
}
