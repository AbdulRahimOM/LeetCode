package medium

import "sort"

func MaxTotalReward(rewardValues []int) int {
	sort.Ints(rewardValues)
	length := len(rewardValues)

	res := [][2]int{{rewardValues[length-1] - 1, rewardValues[length-1]}}
	for i := length - 2; i >= 0; i-- {
		for j := range res {
			if res[j][0] >= rewardValues[i] {
				temp := [2]int{min(res[j][0]-rewardValues[i], rewardValues[i]-1), res[j][1] + rewardValues[i]}
				skip := false
				for k := range res {
					if res[k][0] == temp[0] {
						if res[k][1] < temp[1] {
							res[k] = temp
						}
						//other 2 cases ignored
						skip = true
						break
					} else if res[k][0] > temp[0] { //
						if res[k][1] >= temp[1] {
							//ignore
							skip = true
							break
						}
					} else { //=> res[k][0] < temp[0]	//^
						if res[k][1] <= temp[1] {
							res[k] = temp
							skip = true
							break
						}
					}
				}
				if skip {
					continue
				}
				res = append(res, temp)
			}
		}
	}

	max := 0
	for i := range res {
		if res[i][1] > max {
			max = res[i][1]
		}
	}
	return max
}
