package medium

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
