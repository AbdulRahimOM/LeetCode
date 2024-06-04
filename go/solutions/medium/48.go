package medium

// 48. Rotate Image
// Medium

// You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).
// You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.

// Example 1:
// Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
// Output: [[7,4,1],[8,5,2],[9,6,3]]
// Example 2:

// Input: matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
// Output: [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]

// Constraints:

// n == matrix.length == matrix[i].length
// 1 <= n <= 20
// -1000 <= matrix[i][j] <= 1000


// 1st way
func Rotate1(matrix [][]int) {
	len := int16(len(matrix[0]))
	arr := make([]int, len*len)
	var i, j int16
	for i = 0; i < len; i++ {
		for j = len - 1; j >= 0; j-- {
			arr[i*len+len-1-j] = matrix[j][i]
		}
	}
	for i = 0; i < len; i++ {
		for j = 0; j < len; j++ {
			matrix[i][j] = arr[i*len+j]
		}
	}
}
//2nd way
func Rotate2(matrix [][]int) {
	len := int8(len(matrix[0]))
	arr := make([]int, 4*len-4)
	var i,outer,round int8
	for round=len/2-1;round>=0;round--{
		outer=len-1-round
		for i = outer-round-1; i >=0 ; i-- {
			arr[i] = matrix[outer-i][round]
			arr[outer-round+i] = matrix[round][round+i]
			arr[2*(outer-round)+i] = matrix[round+i][outer]
			arr[3*(outer-round)+i] = matrix[outer][outer-i]
		}
		for i = outer-round-1; i >=0 ; i-- {
			matrix[round][round+i]=arr[i]
			matrix[round+i][outer]=arr[outer-round+i]
			matrix[outer][outer-i]=arr[2*(outer-round)+i]
			matrix[outer-i][round]=arr[3*(outer-round)+i]
		}
	}
}