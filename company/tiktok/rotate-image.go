/*
 * @Author: hongwei.sun
 * @Date: 2024-04-09 23:04:53
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-09 23:05:04
 * @Description: file content
 */
// https://leetcode.com/problems/rotate-image/
/*
You are given an n x n 2D matrix representing an image.

Rotate the image by 90 degrees (clockwise).

Note:

You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.

Example 1:

Given input matrix =
[
  [1,2,3],
  [4,5,6],
  [7,8,9]
],

rotate the input matrix in-place such that it becomes:
[
  [7,4,1],
  [8,5,2],
  [9,6,3]
]
Example 2:

Given input matrix =
[
  [ 5, 1, 9,11],
  [ 2, 4, 8,10],
  [13, 3, 6, 7],
  [15,14,12,16]
],

rotate the input matrix in-place such that it becomes:
[
  [15,13, 2, 5],
  [14, 3, 4, 1],
  [12, 6, 8, 9],
  [16, 7,10,11]
]
*/
/*
旋转90度有以下变化:
1, 原来的row   ---->  be matrix column and index变成了从后往前
2，原来的cloumn ---->  be matrix row index 不变
*/
func rotate(matrix [][]int) {
  n := len(matrix)
  res := make([][]int, n)
  for i := 0; i < n; i++ {
    res[i] = make([]int, n)
  }
  for i := 0; i < n; i++ {
    for j := 0; j < n; j++ {
      res[j][n-1-i] = matrix[i][j]
    }
  }
  // because it requires to be in place
  for i := 0; i < n; i++ {
    for j := 0; j < n; j++ {
      matrix[i][j] = res[i][j]
    }
  }
}
