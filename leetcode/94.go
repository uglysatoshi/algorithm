package leetcode

func inorderTraversal(root *TreeNode) []int {
	var r []int
	var s []*TreeNode
	curr := root
	for curr != nil || len(s) > 0 {
		for curr != nil {
			s = append(s, curr)
			curr = curr.Left
		}
		curr = s[len(s)-1]
		s = s[:len(s)-1]
		r = append(r, curr.Val)
		curr = curr.Right
	}
	return r
}
