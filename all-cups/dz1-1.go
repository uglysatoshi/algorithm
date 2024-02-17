package main

import (
	"fmt"
)

func main() {
	var N, number int
	_, _ = fmt.Scan(&N)
	arr := make([]int, N)
	for i := 0; i < len(arr); i++ {
		_, _ = fmt.Scan(&arr[i])
	}

	counter := 0

	_, _ = fmt.Scan(&number)

	for _, v := range arr {
		if number != v {
			counter++
		}
	}

	fmt.Printf("%d", counter)
}
