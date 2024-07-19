package leetcode

func luckyNumbers(matrix [][]int) []int {
	var r []int

	findMinColumn := func(row int) int {
		minVal, minCol := matrix[row][0], 0
		for j := 1; j < len(matrix[row]); j++ {
			if matrix[row][j] < minVal {
				minVal, minCol = matrix[row][j], j
			}
		}
		return minCol
	}

	isMaxInColumn := func(val, col int) bool {
		for i := 0; i < len(matrix); i++ {
			if matrix[i][col] > val {
				return false
			}
		}
		return true
	}

	for i := 0; i < len(matrix); i++ {
		minCol := findMinColumn(i)
		candidate := matrix[i][minCol]

		if isMaxInColumn(candidate, minCol) {
			r = append(r, candidate)
		}
	}

	return r
}
