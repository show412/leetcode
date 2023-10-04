// https://leetcode.com/problems/maximal-square/
/*
Given a 2D binary matrix filled with 0's and 1's,
find the largest square containing only 1's and return its area.

Example:

Input:

1 0 1 0 0
1 0 1 1 1
1 1 1 1 1
1 0 0 1 0

Output: 4
*/
// it's a typical dp to handle the matrix square
/*
dp[i][j] means the max len with 1 at the right cell of matrix that could have a square
refer to the https://leetcode.com/problems/maximal-square/solution/
*/
// 之所以定义dp 为长宽都加1 就是为了防止各种edge case的情况 详见解法2
func maximalSquare(matrix [][]byte) int {
	maxLen := 0
	// define the dp function
	dp := make([][]int, len(matrix)+1)
	for i := 0; i < len(dp); i++ {
		if len(matrix) > 0 {
			dp[i] = make([]int, len(matrix[0])+1)
		} else {
			dp[i] = make([]int, 1)
		}
	}
	// init the dp function
	dp[0][0] = 0
	// process the dp
	for i := 1; i <= len(matrix); i++ {
		for j := 1; j <= len(matrix[0]); j++ {
			if matrix[i-1][j-1] == '1' {
				// 三个方向最小的正方形的长度
				dp[i][j] = min(dp[i-1][j], dp[i-1][j-1], dp[i][j-1]) + 1
				maxLen = max(maxLen, dp[i][j])
			}
		}
	}
	return maxLen * maxLen
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		} else {
			return c
		}
	} else {
		if b < c {
			return b
		} else {
			return c
		}
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		} else {
			return c
		}
	} else {
		if b < c {
			return b
		} else {
			return c
		}
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 写法2
// 之所以定义dp 为长宽都加1 就是为了防止各种edge case的情况
func maximalSquare(matrix [][]byte) int {
	maxLen := 0
	if len(matrix) == 0 {
		return 0
	}

	if len(matrix[0]) == 1 {
		for i := 0; i < len(matrix); i++ {
			if matrix[i][0] == '1' {
				return 1
			}
		}
		return 0
	}

	if len(matrix) == 1 {
		for i := 0; i < len(matrix[0]); i++ {
			if matrix[0][i] == '1' {
				return 1
			}
		}
		return 0
	}

	// define the dp function
	dp := make([][]int, len(matrix))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(matrix[0]))
	}

	// init the dp function
	if matrix[0][0] == '1' {
		maxLen = 1
		dp[0][0] = 1
		if matrix[0][1] == '1' {
			dp[0][1] = 1
			maxLen = 1
		} else {
			dp[0][1] = 1
		}
		if matrix[1][0] == '1' {
			dp[1][0] = 1
			maxLen = 1
		} else {
			dp[1][0] = 1
		}
	} else {
		dp[0][0] = 0
		if matrix[0][1] == '1' {
			dp[0][1] = 1
			maxLen = 1
		} else {
			dp[0][1] = 0
		}
		if matrix[1][0] == '1' {
			dp[1][0] = 1
			maxLen = 1
		} else {
			dp[1][0] = 0
		}
	}
	// process the dp
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j-1], dp[i][j-1]) + 1
				maxLen = max(maxLen, dp[i][j])
			}
		}
	}
	return maxLen * maxLen
}

func min(a, b, c int) int {
	v := a
	if a < b {
		if a < c {
			v = a

		} else {
			v = c
		}
	} else {
		if b < c {
			v = b
		} else {
			v = c
		}
	}
	return v
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
