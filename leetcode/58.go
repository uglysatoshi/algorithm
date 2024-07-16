package leetcode

func lengthOfLastWord(s string) int {
	r := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			r++
		} else {
			if r > 0 {
				break
			}
		}
	}
	return r
}
