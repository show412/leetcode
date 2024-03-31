// https://leetcode.com/problems/spiral-matrix-ii/
/*
Given a positive integer n, generate a square matrix filled with elements from 1 to n2 in spiral order.

Example:

Input: 3
Output:
[
 [ 1, 2, 3 ],
 [ 8, 9, 4 ],
 [ 7, 6, 5 ]
]
*/
func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	left, right, top, bottom := 0, len(res[0])-1, 0, len(res)-1
	num := 1
	for {
		for i := left; i <= right; i++ {
			res[top][i] = num
			num++
		}
		top++
		if top > bottom {
			break
		}

		for i := top; i <= bottom; i++ {
			res[i][right] = num
			num++
		}
		right--
		if left > right {
			break
		}

		for i := right; i >= left; i-- {
			res[bottom][i] = num
			num++
		}
		bottom--
		if top > bottom {
			break
		}

		for i := bottom; i >= top; i-- {
			res[i][left] = num
			num++
		}
		left++
		if left > right {
			break
		}
	}
	return res
}