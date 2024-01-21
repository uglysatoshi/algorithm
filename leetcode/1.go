package main

import "fmt"

func twoSum(nums []int, target int) []int {
	answer := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i != j && nums[i]+nums[j] == target {
				answer[0] = i
				answer[1] = j
				return answer
			}
		}
	}

	return nil
}

func main() {
	// testcases
	fmt.Printf("%v\n", twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Printf("%v\n", twoSum([]int{3, 2, 4}, 6))
	fmt.Printf("%v\n", twoSum([]int{3, 3}, 6))
}
