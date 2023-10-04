// https://leetcode.com/problems/search-a-2d-matrix-ii/
/*
Write an efficient algorithm that searches for a value in an m x n matrix.
This matrix has the following properties:

Integers in each row are sorted in ascending from left to right.
Integers in each column are sorted in ascending from top to bottom.
Example:

Consider the following matrix:

[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
Given target = 5, return true.

Given target = 20, return false.
*/
/*
找到左下角的值， 当小于target的时候 列增加 当大于target的时候 行减少
because the rows are sorted from left-to-right,
we know that every value to the right of the current value is larger.
Therefore, if the current value is already larger than target,
we know that every value to its right will also be too large.
A very similar argument can be made for the columns,
so this manner of search will always find target in the matrix (if it is present).
*/
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	if len(matrix[0]) == 0 {
		return false
	}
	r := len(matrix) - 1
	c := 0
	// count := 0

	for r >= 0 && c < len(matrix[0]) {
		// fmt.Println(matrix[r][c])
		if target == matrix[r][c] {
			return true
		}
		if target < matrix[r][c] {
			r--
		} else {
			c++
		}
	}
	return false
}
