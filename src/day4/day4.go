package day4

import (
	"mikesigs/aoc-2021/src/day4/bingo"
	"mikesigs/aoc-2021/src/shared"
	"strings"
)

func Part1() int {
	lines, err := shared.ReadLines("/workspace/aoc-2021/data/day4.txt")
	shared.Check(err)

	games := buildGames(lines)
	draws := strings.Split(lines[0], ",")

	for _, draw := range draws {
		for i := 0; i < len(games); i++ {
			game := &games[i]
			if game.Mark(draw) {
				// fmt.Print(game.String())
				if game.IsWinner() {
					return game.WinningScore
				}
			}
		}
	}

	return 0
}

func Part2() int {
	lines, err := shared.ReadLines("/workspace/aoc-2021/data/day4.txt")
	shared.Check(err)

	games := buildGames(lines)
	draws := strings.Split(lines[0], ",")

	var winningScore int
	for _, draw := range draws {
		// fmt.Printf("Draw: %s\n", draw)
		for i := 0; i < len(games); i++ {
			game := &games[i]
			if !game.IsWinner() && game.Mark(draw) {
				// fmt.Print(game.String())
				if game.IsWinner() {
					winningScore = game.WinningScore
				}
			}
		}
	}

	return winningScore
}

func buildGames(lines []string) []bingo.Bingo {
	var games []bingo.Bingo
	var grid [5][5]string
	for b := 2; b < len(lines); b += 6 {
		for i, line := range lines[b : b+6] {
			for j, number := range strings.Fields(line) {
				grid[i][j] = number
			}
		}
		games = append(games, bingo.Bingo{ID: len(games) + 1, Game: grid})
	}
	return games
}
