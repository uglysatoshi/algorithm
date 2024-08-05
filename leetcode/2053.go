package leetcode

func kthDistinct(arr []string, k int) string {
	m := make(map[string]int)
	var order []string
	for _, v := range arr {
		if m[v] == 0 {
			order = append(order, v)
		}
		m[v]++
	}
	cnt := 0
	for _, v := range order {
		if m[v] == 1 {
			cnt++
			if cnt == k {
				return v
			}
		}
	}
	return ""
}
