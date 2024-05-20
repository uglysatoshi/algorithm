package codewars

func GetSum(a, b int) int {
	var sum int

	if a > b {
		for i := b; i <= a; i++ {
			sum += i
		}
	} else {
		for i := a; i <= b; i++ {
			sum += i
		}
	}

	return sum
}
