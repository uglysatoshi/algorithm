package leetcode

import "math"

func ScoreOfString(s string) int {
    res := 0
    for i := 0; i < len(s)-1; i++ {

        res += int(math.Abs(float64(s[i]) - float64(s[i+1])))
    }
    return res
}
