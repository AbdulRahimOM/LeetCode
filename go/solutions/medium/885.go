package medium

// Runtime: 12 ms, Memory: 6.4 MB (beats 26.67%, 80.00%)
func SpiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
	result := make([][]int, 0, rows*cols)
	r := rStart
	c := cStart
	countRemaining := int16(rows * cols)
	i := 0
	for {
		if c >= 0 {
			for r = rStart + i; r >= rStart-i; r-- {
				if r >= rows {
					r = rows - 1
				} else if r < 0 {
					break
				}
				result = append(result, []int{r, c})
				countRemaining--
				if countRemaining == 0 {
					return result
				}
			}
		}
		r = rStart - i
		if r >= 0 {
			for c = cStart + 1 - i; c <= cStart+1+i; c++ {
				if c < 0 {
					c = 0
				} else if c >= cols {
					break
				}
				result = append(result, []int{r, c})
				countRemaining--
				if countRemaining == 0 {
					return result
				}
			}
		}
		c = cStart + i + 1
		if c < cols {
			for r = rStart + 1 - i; r <= rStart+1+i; r++ {
				if r < 0 {
					r = 0
				} else if r >= rows {
					break
				}
				result = append(result, []int{r, c})
				countRemaining--
				if countRemaining == 0 {
					return result
				}
			}
		}
		r = rStart + i + 1
		if r < rows {
			for c = cStart + i; c >= cStart-i; c-- {
				if c >= cols {
					c = cols - 1
				} else if c < 0 {
					break
				}
				result = append(result, []int{r, c})
				countRemaining--
				if countRemaining == 0 {
					return result
				}
			}
		}
		c = cStart - i
		r++
		c--
		i++
	}
}
