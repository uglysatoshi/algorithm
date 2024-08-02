package train_summer_2024

import (
	"fmt"
	"math"
)

func RoundingError() {
	var t int
	_, _ = fmt.Scanln(&t)
	for i := 0; i < t; i++ {
		var n int
		var p, result, v float64
		_, _ = fmt.Scanf("%d %f", &n, &p)
		for j := 0; j < n; j++ {
			_, _ = fmt.Scanln(&v)
			result += math.Abs(roundInt(v*p/100) - roundCorrect(v*p/100))
		}
		fmt.Printf("%.2f\n", result)
	}
}
func roundInt(x float64) float64 {
	return math.Trunc(x)
}
func roundCorrect(x float64) float64 {
	return math.Round(x*100) / 100
}
