package leetcode

func IsValid(s string) bool {
	var stack []rune
	for _, ch := range s {
		if ch == '(' || ch == '[' || ch == '{' {
			stack = append(stack, ch)
		} else {
			if len(stack) == 0 {
				return false
			}
			c := stack[len(stack)-1]
			if c == '(' && ch == ')' || c == '[' && ch == ']' || c == '{' && ch == '}' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
