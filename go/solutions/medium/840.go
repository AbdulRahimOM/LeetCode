package medium

func NumMagicSquaresInside(grid [][]int) int {
	magicGridCnt := 0
	for r := 1; r < len(grid)-1; r++ {
		for c := 1; c < len(grid[0])-1; c++ {
			if grid[r][c] == 5 {
				if grid[r][c+1] != 5 &&
					grid[r-1][c-1] > 0 &&
					grid[r-1][c] > 0 &&
					grid[r-1][c+1] > 0 &&
					grid[r][c-1] > 0 &&
					grid[r][c+1] > 0 &&
					grid[r+1][c-1] > 0 &&
					grid[r+1][c] > 0 &&
					grid[r+1][c+1] > 0 &&
					grid[r-1][c-1]+grid[r-1][c]+grid[r-1][c+1] == 15 &&
					grid[r+1][c-1]+grid[r+1][c]+grid[r+1][c+1] == 15 &&
					grid[r-1][c-1]+grid[r][c-1]+grid[r+1][c-1] == 15 &&
					grid[r-1][c+1]+grid[r][c+1]+grid[r+1][c+1] == 15 &&
					grid[r-1][c-1]+grid[r+1][c+1] == 10 &&
					grid[r-1][c]+grid[r+1][c] == 10 {
					// grid[r][c-1]+grid[r][c+1]==10	//reduntant (other equations are sufficient)
					// grid[r-1][c+1]+grid[r+1][c-1]==10 //reduntant (other equations are sufficient)
					magicGridCnt++
					c += 2
				}
			}
		}
	}
	return magicGridCnt
}