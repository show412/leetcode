/*
 * @Author: hongwei.sun
 * @Date: 2024-04-03 14:54:08
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-10 23:01:14
 * @Description: file content
 */
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
 
 


func sumWithForLoop(numbers []int) int {
    sum := 0
    for _, num := range numbers {
        sum += num
    }
    return sum
}

 func max(a, b int) int {
	if a > b {
		return a
	}
	return b
 }

 func min(a, b int) int {
	if a < b {
		return a
	}
	return b
 }
