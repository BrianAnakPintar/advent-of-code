package utils

import (
	"bufio"
	"os"
)

type Vec2 struct {
  X int
  Y int
}

/*
  ReadFile reads a file and outputs a single continuous string
  of the contents of said file.
*/
func ReadFile(filename string) (string, error) {
  data, err := os.ReadFile(filename)

  if err != nil {
    return "", err
  }

  return string(data), nil
}

/*
  ReadFileLines reads a file and returns a slice of strings,
  where each string is a line from the file.
*/ 
func ReadFileLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func AbsDiff(num1 int, num2 int) int {
  if num1 < num2 {
    return num2 - num1;
  }
  return num1 - num2;
}
