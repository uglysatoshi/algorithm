package std

func merge(a []int, b []int) []int {
	var final []int
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}

func MergeSortRecursive(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	first := MergeSortRecursive(arr[:len(arr)/2])
	second := MergeSortRecursive(arr[len(arr)/2:])
	return merge(first, second)
}
