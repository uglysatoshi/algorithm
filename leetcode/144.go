package leetcode

func preorderTraversal(root *TreeNode) []int {
	r := make([]int, 0)
	if nil == root {
		return r
	}
	s := make([]*TreeNode, 0)
	s = append(s, root)
	for len(s) != 0 {
		x := s[len(s)-1]
		s = s[:len(s)-1]
		r = append(r, x.Val)
		if nil != x.Right {
			s = append(s, x.Right)
		}
		if nil != x.Left {
			s = append(s, x.Left)
		}
	}
	return r
}
