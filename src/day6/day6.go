package day6

import (
	"fmt"
	"mikesigs/aoc-2021/src/shared"
	"strconv"
	"strings"
)

func Part1() int {
	lines, err := shared.ReadLines("/workspace/aoc-2021/data/day6.txt")
	shared.Check(err)
	fish := loadFish(lines[0])

	for i := 0; i < 80; i++ {
		fish = generationA(fish)
	}

	return len(fish)
}

type Node struct {
	next  *Node
	value int
}

func Part2() int64 {
	lines, err := shared.ReadLines("/workspace/aoc-2021/data/day6.txt")
	shared.Check(err)
	inputs := strings.Split(lines[0], ",")
	var list *Node
	for i := range inputs {
		n, err := strconv.Atoi(inputs[i])
		shared.Check(err)
		node := Node{next: list, value: n}
		list = &node
	}

	generations := 256
	for generations > 0 {
		generationL(list)
		generations--
	}

	var count int64
	n := list
	for n != nil {
		count++
		n = n.next
	}

	return count
}

func printList(list *Node) {
	n := list
	for n != nil {
		fmt.Printf("%d->", n.value)
		n = n.next
	}
	fmt.Println()
}

func loadFish(input string) []int {
	inputs := strings.Split(input, ",")
	nums := make([]int, len(inputs))
	for i, x := range inputs {
		n, err := strconv.Atoi(x)
		shared.Check(err)
		nums[i] = n
	}

	return nums
}

func generationA(fish []int) []int {
	for i := range fish {
		if fish[i] == 0 {
			fish[i] = 6
			fish = append(fish, 8)
		} else {
			fish[i]--
		}
	}
	return fish
}

func generationL(list *Node) *Node {
	var new int
	n := list
	for {
		if n.value == 0 {
			n.value = 6
			new++
		} else {
			n.value--
		}
		if n.next != nil {
			n = n.next
		} else {
			break
		}
	}
	for new > 0 {
		n.next = &Node{value: 8}
		n = n.next
		new--
	}

	return list
}
