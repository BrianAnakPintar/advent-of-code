package day3

import (
  "aoc/utils"
  "fmt"
)

func Run() {
  str, ok := utils.ReadFile("day3/input.txt")
  if ok != nil {
    fmt.Println("Error reading file.")
    return
  }

  // Lets rawdog a parser.
  // Use a "state machine", though it's fine if we don't follow the strict design pattern.

  // Part One
  fmt.Println("Part One")
  var parser *MulParser = NewMulParser()
  for _, ch := range str {
    parser.ParseVal(ch) 
  }
  fmt.Println(parser.sum)

  // Part Two
  fmt.Println("Part Two")
  var p2_parser *MulParserTwo = NewMulParserTwo()
  for _, ch := range str {
    p2_parser.ParseVal(ch) 
  }
  fmt.Println(p2_parser.sum)

}
