package hard

//use command below:
//go test -run '^TestBuildMatrix$' -v

import (
	"fmt"
	"testing"
)

const (
	errorMark = "!!!!!!!!!!!!!!!!!!!!!"
	testBegin = "====================================="
)

// TestBuildMatrix tests the BuildMatrix function
func TestBuildMatrix(t *testing.T) {
	tests := []struct {
		k             int
		rowConditions [][]int
		colConditions [][]int
		expected      [][]int
	}{

		// Input: k = 3, rowConditions = [[1,2],[3,2]], colConditions = [[2,1],[3,2]]
		// Output: [[3,0,0],[0,0,1],[0,2,0]]

		// //Input: k = 3, rowConditions = [[1,2],[2,3],[3,1],[2,3]], colConditions = [[2,1]]
		// logger(3, [][]int{{1, 2}, {2, 3}, {3, 1}, {2, 3}}, [][]int{{2, 1}}, true)
		{
			k:             3,
			rowConditions: [][]int{{1, 2}, {2, 3}},
			colConditions: [][]int{{1, 2}, {2, 3}},
			expected:      [][]int{{3, 0, 0}, {0, 0, 1}, {0, 2, 0}},
		},
		// 	Input: k = 3, rowConditions = [[1,2],[2,3],[3,1],[2,3]], colConditions = [[2,1]]
		// Output: []
		{
			k:             3,
			rowConditions: [][]int{{1, 2}, {2, 3}, {3, 1}, {2, 3}},
			colConditions: [][]int{{2, 1}},
			expected:      [][]int{},
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		fmt.Println(testBegin)
		result := BuildMatrix(tt.k, tt.rowConditions, tt.colConditions)
		PrintMatrix(result)
		if len(tt.expected) == 0 {
			if len(result) == 0 {
				t.Logf("Success. Empty matrix")
			} else {
				t.Error("Matrix is not empty. Expected empty matrix", errorMark)
			}
			continue
		}
		//there may be multiple correct answers, so not relying on exact match
		sum := 0
		count := 0
		rowMap := make(map[int]int)
		colMap := make(map[int]int)
		for row := range result {
			for col, num := range result[row] {
				if num != 0 {
					count++
					sum += num
					rowMap[num] = row
					colMap[num] = col
				}
			}
		}
		if sum != tt.k*(tt.k+1)/2 {
			t.Error("Sum is not equal to k*(k+1)/2 (=>doesn't contain all numbers from 1 to k)", errorMark)
			return
		} else if count != tt.k {
			t.Error("Count is not equal to k", errorMark)
			return
		}
		for _, rowCondition := range tt.rowConditions {
			above := rowCondition[0]
			below := rowCondition[1]
			if rowMap[above] >= rowMap[below] {
				t.Error("Row condition failed for [", above, ",", below, "]", errorMark)
				return
			}
		}
		for _, colCondition := range tt.colConditions {
			left := colCondition[0]
			right := colCondition[1]
			if colMap[left] >= colMap[right] {
				t.Error("Col condition failed for [", left, ",", right, "]", errorMark)
				return
			}
		}
		t.Log("Success")
	}
}

func PrintMatrix(m [][]int) {
	for _, r := range m {
		fmt.Println(r)
	}
}
