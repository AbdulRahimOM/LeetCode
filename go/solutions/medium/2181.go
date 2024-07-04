package medium


// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeNodes(head *ListNode) *ListNode {
	cHead := head
	for cHead.Next != nil {
		current := cHead.Next.Next
		for current.Val != 0 {
			cHead.Next.Val += current.Val
			current = current.Next
		}
		cHead.Next.Next = current.Next
		cHead = cHead.Next
	}

	return head.Next
}

//.
//+++++++++++++++++++++++++++++++++++++ Other tries for performace improvement +++++++++++++++++++++++++++++++++++++++++++
func MergeNodes_3(head *ListNode) *ListNode {
	cHead := head
	for cHead.Next != nil {

		current := cHead.Next
		sum := 0
		for current.Val != 0 {
			sum += current.Val
			current = current.Next
		}
		cHead.Next.Val = sum
		cHead.Next.Next = current.Next
		cHead = cHead.Next
	}
	return head.Next
}

func MergeNodes_1(head *ListNode) *ListNode {
	current := head.Next.Next
	for current.Val != 0 {
		head.Next.Val += current.Val
		current = current.Next
	}
	head.Next.Next = current.Next

	cHead := head.Next
	for cHead.Next != nil {

		current = cHead.Next.Next
		for current.Val != 0 {
			cHead.Next.Val += current.Val
			current = current.Next
		}
		cHead.Next.Next = current.Next
		cHead = cHead.Next
	}

	return head.Next
}


//.
//=========================other helper functions made for testing==================================