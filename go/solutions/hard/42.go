package hard

//Trapping Rain Water
//Status at submission: 
//took 10 ms and used memory (5.29 MB)
//Runtime beats 58.24% of golang submission
//Memory usage beats 100% of golang submission

func Trap(height []int) int {
	l,r:=0,len(height)-1
	if r<2{
		return 0
	}
	var totalVolume int
	var peakL int=height[0]
	var peakR int=height[r]

	for l < r {
		if peakL <= peakR {
			for l < r {
				if height[l] > peakL {
					peakL = height[l]
					break
				} else {
					totalVolume += peakL - height[l]
					l++
				}
			}
		} else {
			for l < r {
				if height[r] > peakR {
					peakR = height[r]
					break
				} else {
					totalVolume += peakR - height[r]
					r--
				}
			}
		}
	}
	return totalVolume
}

//took too long (1121 ms) and used memory (6.2 MB)
func Trap_1stAttempt(height []int) int {
	prevPeak := height[0]
	potentialMap := make(map[int]int)
	volume := 0
	for _, h := range height {
		if h >= prevPeak {
			// sum all the potential volume
			for hh, count := range potentialMap {
				volume += count * (prevPeak - hh)
			}
			//reset map
			potentialMap = make(map[int]int)
			prevPeak = h
		} else {
			total := 0
			// if greater than min
			//sum all the potential volume until h
			for hh, count := range potentialMap {
				if hh < h {
					volume += count * (h - hh)
					delete(potentialMap, hh)
					total += count
				}
			}
			//reset map until h and add all count to key h
			//add the current height to the map
			potentialMap[h] = potentialMap[h] + total + 1
		}
	}
	return volume
}