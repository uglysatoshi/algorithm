package leetcode

import (
	"math"
	"slices"
)

func rangeSum(nums []int, n int, left int, right int) int {
	mod := int(math.Pow(10, 9) + 7)
	var a []int
	for i := range len(nums) {
		current := 0
		for j := i; j < len(nums); j++ {
			current += nums[j]
			a = append(a, current)
		}
	}
	slices.Sort(a)
	s := 0
	left--
	for i := left; i < right; i++ {
		s += a[i]
	}
	return s % mod
}
