package leetcode

import (
	"sort"
)

func intersect(n1 []int, n2 []int) []int {
	sort.Ints(n1)
	sort.Ints(n2)
	c1 := 0
	c2 := 0
	var r []int
	for c1 < len(n1) && c2 < len(n2) {
		if n1[c1] < n2[c2] {
			c1++
		} else if n1[c1] > n2[c2] {
			c2++
		} else {
			r = append(r, n1[c1])
			c1++
			c2++
		}
	}
	return r
}
