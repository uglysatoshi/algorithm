package main

import (
    "algorithm/leetcode"
    "fmt"
)

func main() {
    nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

    fmt.Println(nums[:leetcode.RemoveDuplicates(nums)])
}
