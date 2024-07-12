package leetcode

func middleNode(head *ListNode) *ListNode {
	cnt := 0
	for cur := head; cur != nil; cur = cur.Next {
		cnt++
	}
	nCnt := 0
	for cur := head; cur != nil; cur = cur.Next {
		if nCnt == (cnt / 2) {
			return cur
		} else {
			nCnt++
		}
	}
	return nil
}
