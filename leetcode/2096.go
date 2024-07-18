package leetcode

import "strings"

func getDirections(root *TreeNode, startValue int, destValue int) string {
	s := dfs2096(root, []byte{}, startValue)
	d := dfs2096(root, []byte{}, destValue)
	i := 0
	for len(s) > i && len(d) > i && s[i] == d[i] {
		i++
	}
	pU := strings.Repeat("U", len(s)-i)
	pD := string(d[i:])
	return pU + pD
}

func dfs2096(r *TreeNode, a []byte, t int) []byte {
	if r.Val == t {
		return a
	}
	if r.Left != nil {
		n := append(a, 'L')
		if f := dfs2096(r.Left, n, t); f != nil {
			return f
		}
	}
	if r.Right != nil {
		n := append(a, 'R')
		if f := dfs2096(r.Right, n, t); f != nil {
			return f
		}
	}
	return nil
}
