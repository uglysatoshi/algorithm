package leetcode

import "strconv"

func countSeniors(details []string) int {
	r := 0
	for _, v := range details {
		t, _ := strconv.Atoi(v[6:8])
		if t >= 60 {
			r++
		}
	}
	return r
}
