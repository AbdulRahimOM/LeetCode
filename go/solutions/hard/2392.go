package hard

func BuildMatrix(k int, rowConditions [][]int, colConditions [][]int) [][]int {
	rowMap := make(map[int]int,k)
	colMap := make(map[int]int,k)
	for num :=1;num<=k;num++ {
		rowMap[num] = 0
		colMap[num] = 0
	}
	var changed bool
	for range k {
		changed = false
		for i := range rowConditions {
			below:= rowConditions[i][1]
			rowOfAbove := rowMap[rowConditions[i][0]]
			if rowOfAbove >= rowMap[below] {
				changed = true
				rowMap[below] = rowOfAbove + 1
			}
		}
		if !changed {
			break
		}
	}
	if changed { //cycle limit reached
		return [][]int{}
	}

	for range k {
		changed = false
		for i := range colConditions {
			colOfLeft := colMap[colConditions[i][0]]
			right := colConditions[i][1]
			if colOfLeft >= colMap[right] {
				colMap[right] = colOfLeft + 1
				changed = true
			}
		}
		if !changed {
			break
		}
	}
	if changed { //cycle limit reached
		return [][]int{}
	}

	rowContent := make([][]int, k)
	for num, row := range rowMap {
		rowContent[row] = append(rowContent[row], num)
	}

	r := 0
	for i := range rowContent {
		for _, num := range rowContent[i] {
			rowMap[num] = r
			r++
		}
	}

	colContent := make([][]int, k)
	for num, col := range colMap {
		colContent[col] = append(colContent[col], num)
	}

	c := 0
	for i := range colContent {
		for _, num := range colContent[i] {
			colMap[num] = c
			c++
		}
	}

	// Create the matrix
	matrix := make([][]int, k)
	for i := range matrix {
		matrix[i] = make([]int, k)
	}
	for i := range k {
		matrix[rowMap[i+1]][colMap[i+1]] = i + 1
	}
	return matrix
}
