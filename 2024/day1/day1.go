package day1

import (
	"aoc/utils"
	"container/heap"
	"fmt"
	"strconv"
	"strings"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func Run() {
	fmt.Println("AOC - Day 1")
	lines, err := utils.ReadFileLines("day1/input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	h1 := &IntHeap{}
	h2 := &IntHeap{}
	heap.Init(h1)
	heap.Init(h2)

	n := 0
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) != 2 {
			fmt.Println("Error: line does not contain exactly 2 numbers:", line)
			continue
		}

		// Convert the parts to integers
		num1, err1 := strconv.Atoi(parts[0])
		if err1 != nil {
			fmt.Println("Error converting first number:", err1)
			return
		}
		heap.Push(h1, num1)

		num2, err2 := strconv.Atoi(parts[1])
		if err2 != nil {
			fmt.Println("Error converting second number:", err2)
			return
		}
		heap.Push(h2, num2)
		n++
	}

	sum := 0
	for i := 0; i < n; i++ {
		num1 := heap.Pop(h1).(int)
		num2 := heap.Pop(h2).(int)
		diff := float64(num1 - num2)
		if diff < 0 {
			diff *= -1
		}
		sum += int(diff)
	}
	fmt.Println(sum)

	PartTwo()
}

func PartTwo() {
	lines, err := utils.ReadFileLines("day1/input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var list_one []int
	occurences := make(map[int]int, 1000)
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) != 2 {
			fmt.Println("Error: line does not contain exactly 2 numbers:", line)
			continue
		}

		// Convert the parts to integers
		num1, err1 := strconv.Atoi(parts[0])
		if err1 != nil {
			fmt.Println("Error converting first number:", err1)
			return
		}

		num2, err2 := strconv.Atoi(parts[1])
		if err2 != nil {
			fmt.Println("Error converting second number:", err2)
			return
		}

		list_one = append(list_one, num1)
		occurences[num2]++
	}

	similarity_score := 0
	for _, val := range list_one {
		similarity_score += val * occurences[val]
	}
	fmt.Println("PART TWO")
	fmt.Println(similarity_score)
}
