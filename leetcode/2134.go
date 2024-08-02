package leetcode

func minSwaps(nums []int) int {
	ones := count(nums)
	if ones == 0 {
		return 0
	}
	result := ones
	onesAtWindow := 0

	for right := 0; right < ones-1; right++ {
		if nums[right] == 1 {
			onesAtWindow++
		}
	}
	for left := 0; left < len(nums); left++ {
		if left-1 >= 0 && nums[left-1] == 1 {
			onesAtWindow--
		}
		if nums[(left+ones-1)%len(nums)] == 1 {
			onesAtWindow++
		}
		result = min(result, ones-onesAtWindow)
	}
	return result
}

func count(a []int) int {
	result := 0
	for _, v := range a {
		if v == 1 {
			result++
		}
	}
	return result
}
