package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateList(nums []int) *ListNode {
	length := len(nums)
	if length == 0 {
		return &ListNode{}
	}
	if length == 1 {
		return &ListNode{
			Val: nums[0],
		}
	}

	head := &ListNode{
		Val: nums[0],
		Next: &ListNode{
			Val: nums[1],
		},
	}
	current := head
	for i := 2; i < length; i++ {
		current.Next.Next = &ListNode{
			Val: nums[i],
		}
		current = current.Next
	}
	return head
}
func (head *ListNode)DisplayList() {
	fmt.Print("List= ")
	if head == nil {
		fmt.Println("nil")
		return
	} else {
		fmt.Print(head.Val)
	}
	current := head.Next
	for current != nil {
		fmt.Print(",", current.Val)
		current = current.Next
	}
	fmt.Println()
}
