package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (TreeNode *TreeNode) PrintTree() {
	if TreeNode != nil {
		TreeNode.printThrough("")
	}
}

func ArrayToTree(arr []interface{}) *TreeNode {
	if len(arr) == 0 {
		return nil
	}

	// Create the root node
	root := &TreeNode{Val: arr[0].(int)}
	queue := []*TreeNode{root}
	i := 1

	for i < len(arr) {
		// Get the current node from the front of the queue
		current := queue[0]
		queue = queue[1:]

		// Assign the left child if not nil
		if i < len(arr) && arr[i] != nil {
			current.Left = &TreeNode{Val: arr[i].(int)}
			queue = append(queue, current.Left)
		}
		i++

		// Assign the right child if not nil
		if i < len(arr) && arr[i] != nil {
			current.Right = &TreeNode{Val: arr[i].(int)}
			queue = append(queue, current.Right)
		}
		i++
	}

	return root
}

func (n *TreeNode) printThrough(preText string) {
	strBlank := "   "
	strT := "├──"
	strL := "└──"
	strLine := "│   "

	fmt.Println(n.Val)
	if n.Left != nil && n.Right != nil {

		fmt.Print(preText, strT)
		n.Right.printThrough(preText + strLine)

		fmt.Print(preText, strL)
		n.Left.printThrough(preText + strBlank)

	} else if n.Right != nil {
		fmt.Print(preText, strL)
		n.Right.printThrough(preText + strBlank)
	} else if n.Left != nil {
		fmt.Print(preText, strL)
		n.Left.printThrough(preText + strBlank)
	}
}