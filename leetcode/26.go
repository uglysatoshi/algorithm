package leetcode

func RemoveDuplicates(nums []int) int {
    previous := nums[0]
    counter := 1

    for i := range nums {
        if nums[i] != previous {
            nums[counter] = nums[i]
            counter++
        }
        previous = nums[i]
    }

    return counter
}
