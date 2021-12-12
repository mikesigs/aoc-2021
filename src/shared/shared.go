package shared

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadLines(dataFile string) ([]string, error) {
	file, err := os.Open("/workspace/aoc-2021/data/" + dataFile)
	Check(err)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func SplitInts(input string, sep string) []int {
	inputs := strings.Split(input, sep)
	nums := make([]int, len(inputs))
	for i, x := range inputs {
		n, err := strconv.Atoi(x)
		Check(err)
		nums[i] = n
	}

	return nums
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func Abs(input int) int {
	if input < 0 {
		return -input
	}
	return input
}
