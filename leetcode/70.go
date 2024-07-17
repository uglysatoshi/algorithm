package leetcode

func climbStairs(n int) int {
	x := 0
	y := 1
	for i := 0; i < n; i++ {
		a := y
		y += x
		x = a
	}
	return y
}
