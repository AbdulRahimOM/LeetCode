package easy

// More performance optimized version than the initial version (below this function)
// Runtime: 8 ms, beats 97.22% (At time of submission)
// Memory Usage: 5.51 MB, beats 27.78% (At time of submission)
func LuckyNumbers(matrix [][]int) []int {
	result := []int{}
	maxes := make([]int, len(matrix[0]))
	minIndexes := make([]int, len(matrix))
	copy(maxes, matrix[0])
	for i, row := range matrix {
		minIndex := 0
		minValue := row[0]
		for j := range row {
			if row[j] < minValue {
				minValue = row[j]
				minIndex = j
			}
			if row[j] > maxes[j] {
				maxes[j] = row[j]
			}
		}
		minIndexes[i] = minIndex
	}
	for i, c := range minIndexes {
		if maxes[c] == matrix[i][c] {
			result = append(result, maxes[c])
		}
	}
	return result
}

// Initial version
// Runtime: 12 ms, beats 72.22% (At time of submission)
// Memory Usage: 5.46 MB, beats 66.67%
func LuckyNumbers1(matrix [][]int) []int {
	result := []int{}
	for r, row := range matrix {
		min := row[0]
		colIndex := 0
		for i, elem := range row {
			if elem < min {
				min = elem
				colIndex = i
			}
		}
		max := min
		isLuckyNumber := true
		for i := range len(matrix) {
			if matrix[i][colIndex] > max && r != i {
				isLuckyNumber = false
				break
			}
		}
		if isLuckyNumber {
			result = append(result, min)
		}
	}
	return result
}
