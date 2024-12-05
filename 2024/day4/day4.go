package day4

import (
	"aoc/utils"
	"fmt"
)

type Direction uint

const (
	Up Direction = iota
	Down
	Left
	Right

	// Diagonals.
	TopLeft
	TopRight
	BotLeft
	BotRight
)

type Vec2 struct {
	X int
	Y int
}

func NewPos(pos Vec2, dir Direction) Vec2 {
	switch dir {
	case Up:
		pos.Y--
	case Down:
		pos.Y++
	case Left:
		pos.X--
	case Right:
		pos.X++
	case TopLeft:
		pos.X--
		pos.Y--
	case TopRight:
		pos.X++
		pos.Y--
	case BotLeft:
		pos.X--
		pos.Y++
	case BotRight:
		pos.X++
		pos.Y++
	}
	return pos
}

var phases map[rune]rune = map[rune]rune{
	'X': 'M',
	'M': 'A',
	'A': 'S',
}

func search_xmas(grid [][]rune, pos Vec2, phase rune, dir Direction) int {
	i, j := pos.Y, pos.X
	m, n := len(grid), len(grid[0])

	// Bounds check.
	if i >= m || i < 0 || j >= n || j < 0 || grid[i][j] != phase {
		return 0
	}

	// Phase == grid[i][j]
	if phase == 'S' {
		return 1
	}

	newPos := NewPos(pos, dir)
	nextPhase, _ := phases[phase]
	return search_xmas(grid, newPos, nextPhase, dir)
}

func PartOne(grid [][]rune) int {
	m, n := len(grid), len(grid[0])

	sum := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 'X' {
				sum += search_xmas(grid, Vec2{X: j, Y: i}, 'X', Up)
				sum += search_xmas(grid, Vec2{X: j, Y: i}, 'X', Down)
				sum += search_xmas(grid, Vec2{X: j, Y: i}, 'X', Left)
				sum += search_xmas(grid, Vec2{X: j, Y: i}, 'X', Right)

				sum += search_xmas(grid, Vec2{X: j, Y: i}, 'X', TopLeft)
				sum += search_xmas(grid, Vec2{X: j, Y: i}, 'X', TopRight)
				sum += search_xmas(grid, Vec2{X: j, Y: i}, 'X', BotRight)
				sum += search_xmas(grid, Vec2{X: j, Y: i}, 'X', BotLeft)
			}
		}
	}

	return sum
}

func CheckLeftMas(subgrid [3][3]rune) bool {
	if subgrid[0][0] == 'M' && subgrid[2][2] == 'S' {
		return true
	} else if subgrid[0][0] == 'S' && subgrid[2][2] == 'M' {
		return true
	}
	return false
}

func CheckRightMas(subgrid [3][3]rune) bool {
	if subgrid[0][2] == 'M' && subgrid[2][0] == 'S' {
		return true
	} else if subgrid[0][2] == 'S' && subgrid[2][0] == 'M' {
		return true
	}

	return false
}

func isX_MAS(subgrid [3][3]rune) bool {
	/*
	  M . S   M . M
	  . A .   . A .
	  M . S   S . S
	*/
	if subgrid[1][1] != 'A' {
		return false
	}
	return CheckLeftMas(subgrid) && CheckRightMas(subgrid)
}

// Function to print a 3x3 subgrid
func printSubgrid(subgrid [3][3]rune) {
	for _, row := range subgrid {
		fmt.Println(string(row[:]))
	}
}

func PartTwo(grid [][]rune) int {
	/* My idea, brute force:
	Parse a 3x3 grid and check if that grid spells MAS as an X
	Go over every 3x3 grid.
	*/
	m, n := len(grid), len(grid[0])
	fmt.Printf("m: %d, n: %d\n", m, n)

	sum := 0
	for i := 0; i <= m-3; i++ {
		for j := 0; j <= n-3; j++ {
			subgrid := [3][3]rune{
				{grid[i][j], grid[i][j+1], grid[i][j+2]},
				{grid[i+1][j], grid[i+1][j+1], grid[i+1][j+2]},
				{grid[i+2][j], grid[i+2][j+1], grid[i+2][j+2]},
			}
			// printSubgrid(subgrid)
			if isX_MAS(subgrid) {
				sum++
			}
		}
	}
	return sum
}

func Run() {
	lines, err := utils.ReadFileLines("day4/input.txt")

	if err != nil {
		panic(err)
	}

	m, _ := len(lines), len(lines[0])

	// Setup our grid.
	var grid [][]rune = make([][]rune, m)
	for i := range grid {
		grid[i] = []rune(lines[i])
	}

	fmt.Println("Part one:")
	fmt.Println(PartOne(grid))

	fmt.Println("Part two:")
	fmt.Println(PartTwo(grid))
}
