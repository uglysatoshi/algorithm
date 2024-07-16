package leetcode

func removeElement(nums []int, val int) int {
	r := 0
	var n []int
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			n = append(n, nums[i])
		} else {
			r++
		}
	}
	for i := 0; i < len(n); i++ {
		nums[i] = n[i]
	}
	return len(nums) - r
}
