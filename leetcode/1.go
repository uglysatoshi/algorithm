package leetcode

func TwoSum(nums []int,

	target int) []int {
	answer := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		for j :=

			0; j < len(nums); j++ {
			if i != j &&
				nums[i]+nums[j] == target {
				answer[0] = i
				answer[1] = j
				return answer
			}
		}
	}
	return nil
}
