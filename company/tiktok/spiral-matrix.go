/*
 * @Author: hongwei.sun
 * @Date: 2024-04-17 10:20:30
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-17 10:20:30
 * @Description: file content
 */
// https://leetcode.com/problems/spiral-matrix/
/*
Given a matrix of m x n elements (m rows, n columns), return all elements of the matrix in spiral order.

Example 1:

Input:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
Output: [1,2,3,6,9,8,7,4,5]
Example 2:

Input:
[
  [1, 2, 3, 4],
  [5, 6, 7, 8],
  [9,10,11,12]
]
Output: [1,2,3,4,8,12,11,10,9,5,6,7]
*/
// 解决这个问题的方法很多 直接的 straitforwad 就可以 也可以直接递归 参考 ruby 的方法
func spiralOrder(matrix [][]int) []int {
	if matrix == nil || len(matrix) == 0 {
		return []int{}
	}

	result := make([]int, 0)
	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1

	for {
		// right
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++
		if top > bottom {
			break
		}

		// down
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--
		if left > right {
			break
		}

		// left
		for i := right; i >= left; i-- {
			result = append(result, matrix[bottom][i])
		}
		bottom--
		if top > bottom {
			break
		}

		// up
		for i := bottom; i >= top; i-- {
			result = append(result, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}

	return result
}