package day10

import (
	"fmt"
	"mikesigs/aoc-2021/src/shared"
	"sort"
	"strings"
)

type CompletionError struct {
	IllegalChar rune
}

func (e *CompletionError) Error() string {
	return fmt.Sprintf("Found illegal character: %q", e.IllegalChar)
}

func Part1() int {
	lines, err := shared.ReadLines("day10.txt")
	shared.Check(err)

	invalidChars := map[rune]int{}

	for _, line := range lines {
		_, err := completeString(line)
		if err == nil {
			continue
		}
		if e, ok := err.(*CompletionError); ok {
			invalidChars[e.IllegalChar]++
		}
	}

	// fmt.Println(invalidChars)
	return invalidChars[')']*3 + invalidChars[']']*57 + invalidChars['}']*1197 + invalidChars['>']*25137
}

func Part2() int {
	lines, err := shared.ReadLines("day10.txt")
	shared.Check(err)

	var remainingLines []string
	for _, line := range lines {
		_, err := completeString(line)
		if err == nil {
			remainingLines = append(remainingLines, line)
		}
	}

	scores := make([]int, len(remainingLines))
	for i, line := range remainingLines {
		completionString, err := completeString(line)
		if err != nil {
			panic(err)
		}

		scores[i] = completionScore(completionString)
		// fmt.Printf("%q:%d\n", completionString, scores[i])
	}

	sort.Ints(scores)
	return scores[len(scores)/2]
}

func completeString(line string) (string, error) {
	var openStack []rune
	for _, char := range line {
		if isOpeningChar(char) {
			openStack = append(openStack, char)
		} else if isMatch(openStack[len(openStack)-1], char) {
			openStack = openStack[:len(openStack)-1]
		} else {
			return "", &CompletionError{IllegalChar: char}
		}
	}

	completion := ""
	for i := len(openStack) - 1; i >= 0; i-- {
		completion += string(ChunkMap[openStack[i]])
	}

	return completion, nil
}

var CharScore = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

func completionScore(completion string) (score int) {
	for _, char := range completion {
		score = score*5 + CharScore[char]
	}

	return
}

func isOpeningChar(char rune) bool {
	return strings.ContainsRune("([{<", char)
}

var ChunkMap = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

func isMatch(opener, closer rune) bool {
	return closer == ChunkMap[opener]
}
