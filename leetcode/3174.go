package leetcode

func ClearDigits(s string) string {
    var stack []byte
    for i := range s {
        if s[i] >= '0' && s[i] <= '9' {
            stack = stack[:len(stack)-1]
        } else {
            stack = append(stack, s[i])
        }
    }
    return string(stack)
}
