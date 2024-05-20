package allcups

import "fmt"

func moveZerosToEnd(arr []int) []int {
	nonZeros := make([]int, 0)
	zeroCount := 0

	for _, num := range arr {
		if num != 0 {
			nonZeros = append(nonZeros, num)
		} else {
			zeroCount++
		}
	}

	for i := 0; i < zeroCount; i++ {
		nonZeros = append(nonZeros, 0)
	}

	return nonZeros
}

func DZ_1_2() {
	var N int
	_, _ = fmt.Scan(&N)
	arr := make([]int, N)
	for i := 0; i < len(arr); i++ {
		_, _ = fmt.Scan(&arr[i])
	}

	arr = moveZerosToEnd(arr)

	for _, v := range arr {
		fmt.Printf("%d ", v)
	}

}
