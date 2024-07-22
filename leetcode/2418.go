package leetcode

func SortPeople(names []string, heights []int) []string {
	heights, names = MergeSortRecursive(heights, names)
	return names
}

func merge(a []int, b []int, c []string, d []string) ([]int, []string) {
	var finalInt []int
	var finalString []string
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] > b[j] {
			finalInt = append(finalInt, a[i])
			finalString = append(finalString, c[i])
			i++
		} else {
			finalInt = append(finalInt, b[j])
			finalString = append(finalString, d[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		finalInt = append(finalInt, a[i])
		finalString = append(finalString, c[i])

	}
	for ; j < len(b); j++ {
		finalInt = append(finalInt, b[j])
		finalString = append(finalString, d[j])
	}
	return finalInt, finalString
}

func MergeSortRecursive(arrInt []int, arrString []string) ([]int, []string) {
	if len(arrInt) <= 1 {
		return arrInt, arrString
	}
	firstInts, firstStrings := MergeSortRecursive(arrInt[:len(arrInt)/2], arrString[:len(arrString)/2])
	secondInts, secondStrings := MergeSortRecursive(arrInt[len(arrInt)/2:], arrString[len(arrString)/2:])
	return merge(firstInts, secondInts, firstStrings, secondStrings)
}
