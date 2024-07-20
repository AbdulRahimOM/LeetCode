package medium

func RestoreMatrix(rowSum []int, colSum []int) [][]int {
	m:=len(rowSum)
	n:=len(colSum)
	
	matrix:=make([][]int,m)
	for r:=range m{
		matrix[r]=make([]int, n)
		for c:=range n{
			if rowSum[r]==0{
				break
			}
			minn:=minimumOfTwoInts(rowSum[r],colSum[c])
			matrix[r][c]=minn
			rowSum[r]-=minn
			colSum[c]-=minn
		}
	}
    return matrix
}
func minimumOfTwoInts(a, b int) int {
	if a < b {
		return a
	}
	return b
}