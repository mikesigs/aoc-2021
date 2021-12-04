package shared

import (
	"bufio"
	"os"
	"strconv"
)

func ReadNums(path string) []int {
	lines, err := ReadLines(path)
	Check(err)
	var nums []int
	for _, line := range lines {
		num, err := strconv.Atoi(line)
		Check(err)
		nums = append(nums, num)
	}
	return nums
}

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	Check(err)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
