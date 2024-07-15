package leetcode

func createBinaryTree(descriptions [][]int) *TreeNode {
	m := make(map[int]*TreeNode)
	hasParent := make(map[int]bool)
	for _, description := range descriptions {
		parent := description[0]
		child := description[1]
		left := description[2]
		hasParent[child] = true
		curr, exist := m[parent]
		if !exist {
			curr = &TreeNode{
				Val: parent,
			}
			m[parent] = curr
		}
		if left == 1 {
			childNode, exist := m[child]
			if !exist {
				childNode = &TreeNode{
					Val: child,
				}
				m[child] = childNode
			}
			curr.Left = childNode
		} else {
			childNode, exist := m[child]
			if !exist {
				childNode = &TreeNode{
					Val: child,
				}
				m[child] = childNode
			}
			curr.Right = childNode
		}
	}
	var root *TreeNode
	for _, description := range descriptions {
		if !hasParent[description[0]] {
			root = m[description[0]]
		}
	}
	return root
}
