package medium

var res []*TreeNode
var del map[int]bool

func DelNodes(root *TreeNode, to_delete []int) []*TreeNode {
	del = make(map[int]bool)
	res = []*TreeNode{}
	for _, v := range to_delete {
		del[v] = true
	}
	traverse(root)
	if !del[root.Val] {
		res = append(res, root)
	}

	return res
}
func  traverse(n *TreeNode) {
	if n == nil {
		return
	}
	traverse(n.Left)
	traverse(n.Right)

	if del[n.Val] {
		if n.Left != nil && !del[n.Left.Val] {
			res = append(res, n.Left)
		}
		if n.Right != nil && !del[n.Right.Val] {
			res = append(res, n.Right)
		}
	} else {
		if n.Left != nil && del[n.Left.Val] {
			n.Left = nil
		}
		if n.Right != nil && del[n.Right.Val] {
			n.Right = nil
		}
	}
}
