package leetcode

func countWords(words1 []string, words2 []string) int {
	r := 0
	m1 := make(map[string]int)
	m2 := make(map[string]int)
	for _, v := range words1 {
		m1[v]++
	}
	for _, v := range words2 {
		m2[v]++
	}

	if len(m1) > len(m2) {
		for v, _ := range m1 {
			if m2[v] == 1 && m1[v] < 2 {
				r++
			}
		}
	} else {
		for v, _ := range m2 {
			if m1[v] == 1 && m2[v] < 2 {
				r++
			}
		}
	}
	return r
}
