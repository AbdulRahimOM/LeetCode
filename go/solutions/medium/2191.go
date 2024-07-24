package medium

import (
	"fmt"
	"sort"
	"strconv"
	"sync"
)

const divident = 4000

// When submitted: Runtime: 476 ms, Memory: 8.59 MB. (Beats 50%, 100%)
func SortJumbled(mapping []int, nums []int) []int {
	newNums := make([][2]int, len(nums))
	newMapping := make([]byte, 10)
	for i, v := range mapping {
		newMapping[i] = byte(v) + 48
	}
	if len(nums) < 1000 {
		for i, num := range nums {
			newNumBytes := []byte(fmt.Sprint(num))
			for j, digit := range newNumBytes {
				newNumBytes[j] = newMapping[int(digit-'0')]
			}
			newNums[i][0] = num
			newNums[i][1], _ = strconv.Atoi(string(newNumBytes))
		}
	} else {
		wg := sync.WaitGroup{}
		quotient := len(nums) / divident
		for i := 0; i < quotient; i++ {
			//send from i*1000 to (i+1)*1000-1
			wg.Add(1)
			go convert(newMapping, nums[i*divident:(i+1)*divident], newNums[i*divident:(i+1)*divident], &wg)

		}
		//send from quotient*1000 to len(nums)-1
		wg.Add(1)
		go convert(newMapping, nums[quotient*divident:], newNums[quotient*divident:], &wg)
		wg.Wait()

	}

	sort.SliceStable(newNums, func(i, j int) bool {
		return newNums[i][1] < newNums[j][1]
	})

	for i := range newNums {
		nums[i] = newNums[i][0]
	}
	return nums
}
func convert(newMapping []byte, nums []int, newNums [][2]int, wg *sync.WaitGroup) {
	for i, num := range nums {

		newNumBytes := []byte(fmt.Sprint(num))
		for j, digit := range newNumBytes {
			newNumBytes[j] = newMapping[int(digit-'0')]
		}
		newNums[i][0] = num
		newNums[i][1], _ = strconv.Atoi(string(newNumBytes))
	}
	wg.Done()
}
