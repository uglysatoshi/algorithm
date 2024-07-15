package leetcode

func createBinaryTree(descriptions [][]int) *TreeNode {

	m := make(map[int]*TreeNode)
	var children []int
	for _, description := range descriptions {
		parent := description[0]
		child := description[1]
		left := description[2]

		children = append(children, child)
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

	key := -1
	for _, description := range descriptions {
		parent := description[0]
		if !contains(children, parent) {
			key = parent
		}
	}

	return m[key]
}

func contains(slice []int, item int) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
