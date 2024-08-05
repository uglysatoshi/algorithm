package leetcode

import "strings"

func uncommonFromSentences(s1 string, s2 string) []string {
	a1 := strings.Split(s1, " ")
	a2 := strings.Split(s2, " ")
	m := make(map[string]int)
	for _, v := range a1 {
		m[v]++
	}
	for _, v := range a2 {
		m[v]++
	}
	r := make([]string, 0)
	for v, i := range m {
		if i == 1 {
			r = append(r, v)
		}
	}
	return r
}
