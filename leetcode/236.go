package leetcode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	leftVal := lowestCommonAncestor(root.Left, p, q)
	rightVal := lowestCommonAncestor(root.Right, p, q)
	if leftVal != nil && rightVal != nil {
		return root
	}
	if leftVal != nil {
		return leftVal
	} else {
		return rightVal
	}
}
