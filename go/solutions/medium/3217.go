package medium

func ModifiedList(nums []int, head *ListNode) *ListNode {
	presenceMap:=make(map[int]bool)
	for _, v := range nums {
		presenceMap[v] = true
	}

	current := head
	for current.Next != nil {
		if presenceMap[current.Next.Val] {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}

	}

	if presenceMap[head.Val] {
		head = head.Next
	}
	return head

}