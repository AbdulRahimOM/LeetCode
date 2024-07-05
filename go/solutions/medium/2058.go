package medium

// // Definition for singly-linked list.	//Already defined in the package
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }


func NodesBetweenCriticalPoints(head *ListNode) []int {
	current := head
	nodeIndex := 1
	initLocal := 0
	prevLocal := 0
	minDist := -1
	for current.Next != nil && current.Next.Next != nil {
		if (current.Val < current.Next.Val && current.Next.Val > current.Next.Next.Val) ||
			(current.Val > current.Next.Val && current.Next.Val < current.Next.Next.Val) {
			if initLocal == 0 {
				initLocal = nodeIndex
				prevLocal = nodeIndex
			} else if minDist == -1 {
				minDist = nodeIndex - prevLocal
				prevLocal = nodeIndex
			} else {
				minDist = min(minDist, nodeIndex-prevLocal)
				prevLocal = nodeIndex
			}
		}
		current = current.Next
		nodeIndex++
	}
	if minDist == -1 {
		return []int{-1, -1}
	}
	maxDist := prevLocal - initLocal
	return []int{minDist, maxDist}
}