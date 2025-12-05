package day6

import (
	"aoc/utils"
	"fmt"
)

type Direction uint8;

const (
  Up Direction = iota
  Down
  Left
  Right
)

type Guard struct {
  pos utils.Vec2
  dir Direction
}

type GridStatus uint8;
const (
  Visited GridStatus = iota
  Unvisited
  Obstacle
)

type GridBlock struct {
  blockType GridStatus
  visited_dirs []Direction
}

func ChangeDirectionAndGoBack(guard *Guard) {
  switch guard.dir {
  case Up:
    guard.pos.Y++
    guard.dir = Right
  case Down:
    guard.pos.Y--
    guard.dir = Left
  case Left:
    guard.pos.X++
    guard.dir = Up
  case Right:
    guard.pos.X--
    guard.dir = Down
  }
}

func OutOfBounds(grid [][]rune, guard *Guard) bool { 
  i, j := guard.pos.Y, guard.pos.X
  m, n := len(grid), len(grid[0])

  return i < 0 || j < 0 || i >= m || j >= n
}

// Does a walk and marks visited nodes onto the grid.
func GuardWalk(grid [][]rune, guard *Guard) {
  i, j := guard.pos.Y, guard.pos.X

  if OutOfBounds(grid, guard) {
    return 
  }

  if grid[i][j] == '#' {
    // Change direction and go back.
    ChangeDirectionAndGoBack(guard)
  } else {
    grid[i][j] = 'X'  // Mark as visited.
  }

  switch guard.dir {
  case Up:
    guard.pos.Y--
  case Right:
    guard.pos.X++
  case Down:
    guard.pos.Y++
  case Left:
    guard.pos.X--
  }

  GuardWalk(grid, guard)
}

func OutOfBoundsTwo(grid [][]GridBlock, guard *Guard) bool { 
  i, j := guard.pos.Y, guard.pos.X
  m, n := len(grid), len(grid[0])

  return i < 0 || j < 0 || i >= m || j >= n
}

func ComplementDirection(origDir Direction, destDir Direction) bool {
  // true if:
  return (origDir == Up && destDir == Right) ||
         (origDir == Right && destDir == Down) ||
         (origDir == Down && destDir == Left) ||
         (origDir == Left && destDir == Up)
}

// Does a walk and marks visited nodes onto the grid.
func GuardWalkTwo(grid [][]GridBlock, guard *Guard, ways *int) {
  i, j := guard.pos.Y, guard.pos.X

  if OutOfBoundsTwo(grid, guard) {
    return 
  }

  if grid[i][j].blockType == Obstacle {
    // Change direction and go back.
    ChangeDirectionAndGoBack(guard)
  } else if grid[i][j].blockType == Unvisited {
    grid[i][j].blockType = Visited  // Mark as visited.
    grid[i][j].visited_dirs = append(grid[i][j].visited_dirs, guard.dir)

  } else if grid[i][j].blockType == Visited {
    for _, dir := range grid[i][j].visited_dirs {
      if ComplementDirection(guard.dir, dir) {
        fmt.Printf("Located at: %d, %d swap dir\n", i, j)
        *ways++
      }
    }
  }

  switch guard.dir {
  case Up:
    guard.pos.Y--
  case Right:
    guard.pos.X++
  case Down:
    guard.pos.Y++
  case Left:
    guard.pos.X--
  }

  GuardWalkTwo(grid, guard, ways)
}

func PartOne(inputPath string) int {
  lines, _ := utils.ReadFileLines(inputPath)

	m, n := len(lines), len(lines[0])

	// Setup our grid.
	var grid [][]rune = make([][]rune, m)
	for i := range grid {
		grid[i] = []rune(lines[i])
	}

  var guard Guard
  for i := 0; i < m; i++ {
    found := false
    for j := 0; j < n; j++ {
      if grid[i][j] == '^' {
        guard.pos.X = j
        guard.pos.Y = i
        guard.dir = Up

        found = true
        break;
      }
    }
    if found {
      break
    }
  }

  GuardWalk(grid, &guard)
  ans := 0
  for i := 0; i < m; i++ {
    for j := 0; j < n; j++ {
      if grid[i][j] == 'X' {
        ans++
      }
    }
  }


  return ans
}

func PartTwo(inputPath string) int {
  lines, _ := utils.ReadFileLines(inputPath)

	m, n := len(lines), len(lines[0])

	// Setup our grid.
	var grid [][]rune = make([][]rune, m)
	for i := range grid {
		grid[i] = []rune(lines[i])
	}
  var betterGrid [][]GridBlock = make([][]GridBlock, m)
	for i := range betterGrid {
		betterGrid[i] = make([]GridBlock, n)
	}

  var guard Guard
  for i := 0; i < m; i++ {
    for j := 0; j < n; j++ {
      betterGrid[i][j].blockType = Unvisited
      if grid[i][j] == '^' {
        guard.pos.X = j
        guard.pos.Y = i
        guard.dir = Up
      } else if (grid[i][j] == '#') {
        betterGrid[i][j].blockType = Obstacle
      }
    }
  }

  ans := 0
  GuardWalkTwo(betterGrid, &guard, &ans)

  return ans
}

func Run() {
  fmt.Println("Sample Input: ")
  fmt.Println(PartOne("day6/sample.txt"))

  fmt.Println("Main Input: ")
  fmt.Println(PartOne("day6/input.txt")) 

  fmt.Println("Part 2")

  fmt.Println("Sample Input: ")
  fmt.Println(PartTwo("day6/sample.txt"))

  fmt.Println("Main Input: ")
  // fmt.Println(PartTwo("day6/input.txt")) 

}
