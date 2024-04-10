/*
 * @Author: hongwei.sun
 * @Date: 2021-01-22 18:45:51
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-10 23:02:04
 * @Description: file content
 */
// https://leetcode.com/problems/minimum-path-sum/
/*
Given a m x n grid filled with non-negative numbers,
find a path from top left to bottom right
which minimizes the sum of all numbers along its path.

Note: You can only move either down or right at any point in time.

Example:

Input:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
Output: 7
Explanation: Because the path 1→3→1→1→1 minimizes the sum.
*/
// DP solution

func minPathSum(grid [][]int) int {
	r := len(grid)
	c := len(grid[0])
    dp := make([][]int, r)
	for i := 0; i < r; i++ {
		dp[i] = make([]int, c)
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if i == 0 || j == 0 {
				if i == 0 && j > 0 {
					dp[i][j] = grid[i][j] + dp[i][j-1]
				} else if i > 0 && j == 0 {
					dp[i][j] = grid[i][j] + dp[i-1][j]
				} else {
					dp[i][j] = grid[i][j]
				}
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}
		}
	}
	return dp[r-1][c-1]
 }

func minPathSum(grid [][]int) int {
	// dp[i][j] means the min path sum on i and j
	dp := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		dp[i] = make([]int, len(grid[0]))
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 && j != 0 {
				dp[i][j] = grid[i][j] + dp[i][j-1]
			} else if i != 0 && j == 0 {
				dp[i][j] = grid[i][j] + dp[i-1][j]
			} else if i != 0 && j != 0 {
				dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])
			} else {
				dp[i][j] = grid[i][j]
			}
		}
	}
	return dp[len(grid)-1][len(grid[0])-1]
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}