package main

import (
	"fmt"
	"log"
	"sort"
)

func main() {
	var arr = make([]float64, 3)
	for i := 0; i < len(arr); i++ {
		_, err := fmt.Scan(&arr[i])
		if err != nil {
			log.Fatal(err)
		}
	}
	sort.Float64s(arr)
	fmt.Printf("%.0f", arr[1])
}
