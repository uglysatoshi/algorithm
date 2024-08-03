package leetcode

import "reflect"

func canBeEqual(target []int, arr []int) bool {
	m1 := make(map[int]int)
	m2 := make(map[int]int)
	for i := 0; i < len(target); i++ {
		m1[target[i]]++
		m2[arr[i]]++
	}
	return reflect.DeepEqual(m1, m2)
}
