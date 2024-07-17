package leetcode

func delNodes(r *TreeNode, toDelete []int) []*TreeNode {
	var f []*TreeNode
	m := make(map[int]bool)
	for _, v := range toDelete {
		m[v] = true
	}
	if !m[r.Val] {
		f = append(f, r)
	}
	dfs1110(r, m, &f)
	return f
}

func dfs1110(n *TreeNode, d map[int]bool, f *[]*TreeNode) *TreeNode {
	if n == nil {
		return nil
	}
	n.Left = dfs1110(n.Left, d, f)
	n.Right = dfs1110(n.Right, d, f)
	if d[n.Val] {
		if n.Left != nil {
			*f = append(*f, n.Left)
		}
		if n.Right != nil {
			*f = append(*f, n.Right)
		}
		return nil
	}
	return n
}
