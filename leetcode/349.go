package leetcode

func intersection(n1 []int, n2 []int) []int {
	r := make([]int, 0)
	m1 := make(map[int]int)
	m2 := make(map[int]int)
	for _, v := range n1 {
		m1[v]++
	}
	for _, v := range n2 {
		m2[v]++
	}
	for v, _ := range m1 {
		if m2[v] > 0 {
			r = append(r, v)
		}
	}
	return r
}
