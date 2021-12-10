package bingo

import (
	"fmt"
	"mikesigs/aoc-2021/src/shared"
	"strconv"
)

const (
	rows int = 5
	cols int = 5
)

type Bingoer interface {
	Mark(number string) bool
	IsWinner() (bool, int)
	String() string
}

type Bingo struct {
	ID           int
	Game         [rows][cols]string
	WinningScore int
}

func (b *Bingo) Mark(value string) bool {
	if b.WinningScore > 0 {
		return true
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if b.Game[i][j] == value {
				b.Game[i][j] = "*" + b.Game[i][j]
				if b.IsWinner() {
					b.WinningScore = b.calculateScore(value)
				}

				return true
			}
		}
	}
	return false
}

func (b *Bingo) IsWinner() bool {
	if b.WinningScore > 0 {
		return true
	}

	var marked int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if b.Game[i][j][0] == '*' {
				marked++
			}
		}

		if marked == 5 {
			return true
		} else {
			marked = 0
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if b.Game[j][i][0] == '*' {
				marked++
			}
		}

		if marked == 5 {
			return true
		} else {
			marked = 0
		}
	}

	return false
}

func (b *Bingo) calculateScore(value string) int {
	drawNum, err := strconv.Atoi(value)
	shared.Check(err)
	var unmarked int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if b.Game[j][i][0] != '*' {
				number, err := strconv.Atoi(b.Game[j][i])
				shared.Check(err)
				unmarked += number
			}
		}
	}

	return unmarked * drawNum
}

func (b *Bingo) String() string {
	var result string
	result += fmt.Sprintf("Game #: %d\n", b.ID)
	for i := 0; i < rows; i++ {
		for _, value := range b.Game[i][:] {
			if len(value) == 1 {
				value = "0" + value
			}
			result += fmt.Sprintf("%4v", value)
		}
		result += "\n"
	}
	result += fmt.Sprintf("Winner: %t\n", b.IsWinner())
	result += fmt.Sprintf("Score: %d\n\n", b.WinningScore)
	return result
}
