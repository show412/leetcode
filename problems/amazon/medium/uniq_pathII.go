// https://leetcode.com/problems/unique-paths-ii/
/*
A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).

The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).

Now consider if some obstacles are added to the grids. How many unique paths would there be?



An obstacle and empty space is marked as 1 and 0 respectively in the grid.

Note: m and n will be at most 100.

Example 1:

Input:
[
  [0,0,0],
  [0,1,0],
  [0,0,0]
]
Output: 2
Explanation:
There is one obstacle in the middle of the 3x3 grid above.
There are two ways to reach the bottom-right corner:
1. Right -> Right -> Down -> Down
2. Down -> Down -> Right -> Right
*/
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	r := len(obstacleGrid)
	c := len(obstacleGrid[0])
	f := make([][]int, r)
	for i := 0; i < r; i++ {
		f[i] = make([]int, c)
	}
	f[0][0] = 1
	for i := 1; i < r; i++ {
		if obstacleGrid[i][0] == 1 {
			f[i][0] = 0
		} else {
			f[i][0] = f[i-1][0]
		}
	}
	for i := 1; i < c; i++ {
		if obstacleGrid[0][i] == 1 {
			f[0][i] = 0
		} else {
			f[0][i] = f[0][i-1]
		}
	}
	for i := 1; i < r; i++ {
		for j := 1; j < c; j++ {
			if obstacleGrid[i][j] == 1 {
				f[i][j] = 0
			} else {
				f[i][j] = f[i-1][j] + f[i][j-1]
			}
		}
	}
	return f[r-1][c-1]
}
