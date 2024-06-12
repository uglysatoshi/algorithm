package codewars

import (
    "math"
)

func NbYear(p0 int, percent float64, aug int, p int) int {
	result := 0
	var aug1, p2, p3 float64
	
	p2 = float64(p0)
	p3 = float64(p)
	aug1 = float64(aug)

	for p2 < p3  {
		p2 = math.Floor(p2 + (p2*(percent/100) + aug1))
		result ++
	}
	
	return result
	
}