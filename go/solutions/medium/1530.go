package medium

var gDistance int
var gCount int

func CountPairs(root *TreeNode, distance int) int {
	gDistance = distance
	gCount = 0
	traverseForLeafs(root,1)

	return gCount
}

func  traverseForLeafs(n *TreeNode,rank int) map[int]int {
	if n == nil {
		return nil
	}
	if n.Left == nil && n.Right == nil {
		return map[int]int{rank: 1}
	}
	map1 := traverseForLeafs(n.Left,rank + 1)
	delete(map1,rank+gDistance)
	map2 := traverseForLeafs(n.Right,rank + 1)
	delete(map2,rank+gDistance)

	if map1==nil{
		return map2
	}
	if map2==nil{
		return map1
	}
	for map1Rank, v1 := range map1 {
		requiredMaxRank := 2*rank + gDistance - map1Rank //rank+gDistance-(map1Rank-rank)
		for i:=2;i<= requiredMaxRank;i++ {
			gCount += map2[i] * v1
		}
		map1[map1Rank] += map2[map1Rank]
	}
	for map2Rank,v2:=range map2{
		if _,exists:=map1[map2Rank];!exists{
			map1[map2Rank]=v2
		}
	}
	return map1
}