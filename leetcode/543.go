package leetcode

func diameterOfBinaryTree(root *TreeNode) int {
	ans := 0
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return -1
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		ans = max(ans, 2+left+right)
		return 1 + max(left, right)
	}
	dfs(root)
	return ans
}
