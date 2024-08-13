package hard

var gN int8
var result [][]string
var stringToEdit []byte

func SolveNQueens(n int) [][]string {
	gN = int8(n)
	stringToEdit = make([]byte, n)
	for i := range stringToEdit {
		stringToEdit[i] = '.'
	}
	result = [][]string{}
	solve(0, [][2]int8{})
	return result

}
func solve(row int8, calcs [][2]int8) {
	if row == gN {
		newResult := []string{}
		for _, calc := range calcs {
			stringToEdit[calc[1]] = 'Q'
			newResult = append(newResult, string(stringToEdit))
			stringToEdit[calc[1]] = '.'
		}
		result = append(result, newResult)
		return
	}
	for col := range gN {
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