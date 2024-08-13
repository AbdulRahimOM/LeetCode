package hard

var gn int8
var result2 int

func TotalNQueens(n int) int {
	gn = int8(n)
	result2 = 0
	solver(0, [][2]int8{})
	return result2

}
func solver(row int8, calcs [][2]int8) {
	if row == gn {
		result2++
		return
	}
	for col := range gn {
		allowed := true
		for _, calc := range calcs {
			r := calc[0]
			c := calc[1]
			if c == col || (r-c) == (row-col) || (r+c) == (row+col) {
				allowed = false
				break
			}
		}
		if allowed {
			solve(row+1, append(calcs, [2]int8{row, col}))
		}
	}
}
