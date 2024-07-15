package leetcode

func NumJewelsInStones(jewels string, stones string) int {
	cnt := 0
	for i := 0; i < len(stones); i++ {
		for j := 0; j < len(jewels); j++ {
			if jewels[j] == stones[i] {
				cnt++
			}
		}
	}
	return cnt
}
