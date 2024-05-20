package allcups

import "fmt"

func DZ_3_1() {
	var n int
	_, _ = fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < len(arr); i++ {
		_, _ = fmt.Scan(&arr[i])
	}
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}

	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
}
