package leetcode

func merge88(a []int, m int, b []int, n int) {
	r := make([]int, 0)
	i := 0
	j := 0
	for i < m && j < n {
		if a[i] < b[j] {
			r = append(r, a[i])
			i++
		} else {
			r = append(r, b[j])
			j++
		}
	}
	for ; i < m; i++ {
		r = append(r, a[i])
	}
	for ; j < n; j++ {
		r = append(r, b[j])
	}

	for k, v := range r {
		a[k] = v
	}
}
