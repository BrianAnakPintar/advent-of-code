package day2

import (
	"aoc/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Run() {
	lines, err := utils.ReadFileLines("day2/input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	ans := 0
	for _, line := range lines {
		strNums := strings.Split(line, " ")
		nums := make([]int, len(strNums))

		// Convert each string to an integer
		for i, str := range strNums {
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Printf("Error converting %s to int: %v\n", str, err)
				return
			}
			nums[i] = num
		}

    if isSafe(nums) {
			ans++
		} else {
      // Part two.
      for i := range nums {
        cpy_nums := make([]int, len(nums))
        copy(cpy_nums, nums)
        slices.Delete(cpy_nums, i, i+1)
        cpy_nums = cpy_nums[:len(cpy_nums)-1]
        if isSafe(cpy_nums) {
          ans++
          break
        }
      }
    }
	}

	fmt.Println(ans)
}

func isSafe(nums []int) bool {
	n := len(nums)	

	prev := nums[0]
	decreasing := nums[1] < nums[0]

	for i := 1; i < n; i++ {
		diff := nums[i] - prev
		if (decreasing &&  (diff >= 0 || -diff > 3)) ||
			 (!decreasing && (diff <= 0 || diff > 3)) {
				return false
		}
		prev = nums[i]
	}
	return true
}
