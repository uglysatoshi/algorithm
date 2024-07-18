package leetcode

func countPairs(root *TreeNode, distance int) int {
	ans := 0
	var dfs func(node *TreeNode) []int
	dfs = func(node *TreeNode) []int {
		if node == nil {
			return []int{}
		}
		if node.Left == nil && node.Right == nil {
			return []int{0}
		}
		leftDistances := dfs(node.Left)
		rightDistances := dfs(node.Right)
		for _, l := range leftDistances {
			for _, r := range rightDistances {
				if l+r+2 <= distance {
					ans++
				}
			}
		}
		var distances []int
		for _, l := range leftDistances {
			distances = append(distances, l+1)
		}
		for _, r := range rightDistances {
			distances = append(distances, r+1)
		}
		return distances
	}

	dfs(root)
	return ans
}
