package day5

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

func CheckValid(nums []int, rules map[int][]int) bool {
  dont_allow_set := map[int]bool{}

  for _, num := range nums {
    _, hasVal := dont_allow_set[num]
    if hasVal == true {
      return false
    }

    // Add all nodes we are not allowed to visit.
    rules_for_num := rules[num]
    for _, num := range rules_for_num {
      dont_allow_set[num] = true
    }
  }
  return true
}

func CreateRules(rules_str []string) map[int][]int {
  rules := make(map[int][]int)

  for _, rule_str := range rules_str {
    keys := strings.Split(rule_str, "|")
    orig, _ := strconv.Atoi(keys[0])
    dest, _ := strconv.Atoi(keys[1])

    rules[dest] = append(rules[dest], orig)
  }

  return rules
}

func GetMiddle(nums []int) int {
  n := len(nums)
  return nums[n/2]
}

func PartOne(pages []string, rules map[int][]int) int {
  res := 0
  
  for _, page := range pages {
    nums_str := strings.Split(page, ",")
    nums := make([]int, len(nums_str))

    for i, num_str := range nums_str {
      nums[i], _ = strconv.Atoi(num_str)
    }

    if CheckValid(nums, rules) {
      res += GetMiddle(nums)
    }
  }

  return res
}

func PartTwo(pages []string, rules map[int][]int) int {
  res := 0
  
  for _, page := range pages {
    nums_str := strings.Split(page, ",")
    nums := make([]int, len(nums_str))

    for i, num_str := range nums_str {
      nums[i], _ = strconv.Atoi(num_str)
    }

    if !CheckValid(nums, rules) {
      sorted, err := topologicalSortForSlice(rules, nums)
      if err != nil {
        fmt.Println("ERROR")
      }
      res += GetMiddle(sorted)
    }
  }

  return res
}

func topologicalSortForSlice(graph map[int][]int, input []int) ([]int, error) {
	filteredGraph := make(map[int][]int)
	nodeInSlice := make(map[int]bool)
	for _, node := range input {
		nodeInSlice[node] = true
	}

	for node, dependencies := range graph {
		if nodeInSlice[node] {
			for _, dep := range dependencies {
				if nodeInSlice[dep] {
					filteredGraph[node] = append(filteredGraph[node], dep)
				}
			}
		}
	}

	inDegree := make(map[int]int)
	for node := range filteredGraph {
		if _, exists := inDegree[node]; !exists {
			inDegree[node] = 0
		}
		for _, neighbor := range filteredGraph[node] {
			inDegree[neighbor]++
		}
	}

	var queue []int
	for _, node := range input { // Only consider nodes in input slice
		if inDegree[node] == 0 {
			queue = append(queue, node)
		}
	}

	var result []int

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		result = append(result, current)

		for _, neighbor := range filteredGraph[current] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// Step 5: Preserve order for unrelated nodes
	seen := make(map[int]bool)
	for _, node := range result {
		seen[node] = true
	}
	for _, node := range input {
		if !seen[node] {
			result = append(result, node)
		}
	}

	// Step 6: Check if there's a cycle
	if len(result) != len(input) {
		return nil, fmt.Errorf("graph has a cycle, topological sort not possible")
	}

	return result, nil
}

func PrintRules(rules map[int][]int) {
  fmt.Println("RULES")
  for k, v := range rules {
    fmt.Printf("%d: ", k)
    for _, val := range v {
      fmt.Printf("%d ", val)
    }
    fmt.Println()
  }
  fmt.Println("END RULES")
}

func Run() {
  vec_str, _ := utils.ReadFileLines("day5/rules.txt")
  pages, _ := utils.ReadFileLines("day5/pages.txt")
  rules := CreateRules(vec_str) 

  // PrintRules(rules)

  fmt.Println("Part 1:")
  fmt.Println(PartOne(pages, rules))

  fmt.Println("Part 2:")
  fmt.Println(PartTwo(pages, rules))
}
