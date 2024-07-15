package medium

// When submitted: 342 ms (beats 64%)	10.85 MB (beats 29%)
func CreateBinaryTree(descriptions [][]int) *TreeNode {
	type details struct {
		Node    *TreeNode
		IsChild bool
	}
	nodeMap := make(map[int]details)
	for _, desc := range descriptions {
		var node *TreeNode
		if _, ok := nodeMap[desc[0]]; ok {
			node = nodeMap[desc[0]].Node
		} else {
			node = &TreeNode{Val: desc[0]}
			nodeMap[desc[0]] = details{node, false}
		}
		if desc[2] == 1 {
			if _, ok := nodeMap[desc[1]]; ok {
				node.Left = nodeMap[desc[1]].Node
				nodeMap[desc[1]] = details{node.Left, true}
			} else {
				node.Left = &TreeNode{Val: desc[1]}
				nodeMap[desc[1]] = details{node.Left, true}
			}
		} else {
			if _, ok := nodeMap[desc[1]]; ok {
				node.Right = nodeMap[desc[1]].Node
				nodeMap[desc[1]] = details{node.Right, true}

			} else {
				node.Right = &TreeNode{Val: desc[1]}
				nodeMap[desc[1]] = details{node.Right, true}
			}
		}
	}
	for _, v := range nodeMap {
		if !v.IsChild {
			return v.Node
		}
	}
	return nil //won't happen
}


//369 ms	8.69 MB
// When submitted: 369 ms (beats 44%)	8.69 MB (beats 82%)
func CreateBinaryTree_OtherAttempts1(descriptions [][]int) *TreeNode {	
	type details struct {
		Node    *TreeNode
		IsChild bool
	}
	nodeMap := make(map[int]details)
	for _, desc := range descriptions {
		var node *TreeNode
		if _, ok := nodeMap[desc[0]]; ok {
			node = nodeMap[desc[0]].Node
		} else {
			node = &TreeNode{Val: desc[0]}
			nodeMap[desc[0]] = details{node, false}
		}
		if desc[2] == 1 {
			if _, ok := nodeMap[desc[1]]; ok {
				node.Left = nodeMap[desc[1]].Node
				if !nodeMap[desc[1]].IsChild {
					nodeMap[desc[1]]=details{node.Left, true}
				}
			} else {
				node.Left = &TreeNode{Val: desc[1]}
				nodeMap[desc[1]] = details{node.Left, true}
			}
		} else {
			if _, ok := nodeMap[desc[1]]; ok {
				node.Right = nodeMap[desc[1]].Node
				if !nodeMap[desc[1]].IsChild {
					nodeMap[desc[1]]=details{node.Right, true}
				}
			} else {
				node.Right = &TreeNode{Val: desc[1]}
				nodeMap[desc[1]] = details{node.Right, true}
			}
		}
	}
	for _,v:=range nodeMap{
		if !v.IsChild{
			return v.Node
		}
	}
	return nil	//won't happen
}