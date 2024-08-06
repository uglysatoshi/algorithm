package leetcode

import "sort"

func minimumPushes(w string) int {
	l := make([]int, 26)
	for _, c := range w {
		l[c-'a']++
	}
	sort.Sort(sort.Reverse(sort.IntSlice(l)))
	r := 0
	for i, f := range l {
		if f == 0 {
			break
		}
		r += (i/8 + 1) * f
	}
	return r
}
