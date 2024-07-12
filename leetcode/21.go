package leetcode

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	final := &ListNode{}
	curr := final
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next
	}
	if list1 != nil {
		final.Next = list1
	} else {
		final.Next = list2
	}
	return final.Next
}