package main

import "fmt"

func romanToInt(s string) int {
	m := make(map[rune]int)

	m['I'] = 1
	m['V'] = 5
	m['X'] = 10
	m['L'] = 50
	m['C'] = 100
	m['D'] = 500
	m['M'] = 1000

	chars := []rune(s)

	r := m[chars[len(chars)-1]]

	for i := len(chars) - 2; i >= 0; i-- {
		if m[chars[i]] < m[chars[i+1]] {
			r -= m[chars[i]]
		} else {
			r += m[chars[i]]
		}
	}

	return r
}

func main() {
	fmt.Printf("%v\n", romanToInt("III"))
}
