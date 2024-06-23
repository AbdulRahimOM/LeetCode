package medium


func MinimumArea(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	left, right, top, bottom, i, j := 0, 0, 0, 0, 0, 0

	//find top
	extraBreak := false
	for i = 0; i < rows; i++ {
		for j = 0; j < cols; j++ {
			if grid[i][j] == 1 {
				top = i
				{   //lets take this as initial value for bottom, left and right
					left = j
					right = j
					bottom = i
				}
				extraBreak = true
				break
			}
		}
		if extraBreak {
			break
		}
	}

    //find left
	for i = top + 1; i < rows; i++ {
		for j = 0; j < left; j++ {
			if grid[i][j] == 1 {
				left = j
				break
			}
		}
	}

    //find bottom
	extraBreak = false
	for i = rows - 1; i >= top; i-- {
		for j = cols - 1; j >= 0; j-- {
			if grid[i][j] == 1 {
				bottom = i
				extraBreak = true
				break
			}
		}
		if extraBreak {
			break
		}
	}

    //find right
	for i = bottom;i >= top ;i--{
		for j = cols - 1; j > right; j-- {
			if grid[i][j] == 1 {
				right = max(right, j)
				break
			}
		}
	}

	return (right - left + 1) * (bottom - top + 1)
}