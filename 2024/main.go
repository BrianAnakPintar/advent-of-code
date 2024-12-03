package main

import (
	"aoc/day1"
	"aoc/day2"
	"aoc/day3"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <day>")
		os.Exit(1)
	}

	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid day:", os.Args[1])
		os.Exit(1)
	}

	switch day {
	case 1:
		day1.Run()
	case 2:
		day2.Run()
  case 3:
    day3.Run()
	default:
		fmt.Printf("Day %d is not implemented yet\n", day)
	}
}
