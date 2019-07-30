// https://leetcode.com/problems/range-sum-query-2d-immutable/
/*
Given a 2D matrix matrix, find the sum of the elements inside the rectangle defined by its upper left corner (row1, col1) and lower right corner (row2, col2).

Range Sum Query 2D
The above rectangle (with the red border) is defined by (row1, col1) = (2, 1) and (row2, col2) = (4, 3), which contains sum = 8.

Example:
Given matrix = [
  [3, 0, 1, 4, 2],
  [5, 6, 3, 2, 1],
  [1, 2, 0, 1, 5],
  [4, 1, 0, 1, 7],
  [1, 0, 3, 0, 5]
]

sumRegion(2, 1, 4, 3) -> 8
sumRegion(1, 1, 2, 2) -> 11
sumRegion(1, 2, 2, 4) -> 12
Note:
You may assume that the matrix does not change.
There are many calls to sumRegion function.
You may assume that row1 ≤ row2 and col1 ≤ col2.
*/
/*
The thought is:
+-----+-+-------+     +--------+-----+     +-----+---------+     +-----+--------+
|     | |       |     |        |     |     |     |         |     |     |        |
|     | |       |     |        |     |     |     |         |     |     |        |
+-----+-+       |     +--------+     |     |     |         |     +-----+        |
|     | |       |  =  |              |  +  |     |         |  -  |              |
+-----+-+       |     |              |     +-----+         |     |              |
|               |     |              |     |               |     |              |
|               |     |              |     |               |     |              |
+---------------+     +--------------+     +---------------+     +--------------+

   sums[i][j]      =    sums[i-1][j]    +     sums[i][j-1]    -   sums[i-1][j-1]   +

												matrix[i-1][j-1]

So, we use the same idea to find the specific area's sum.

+---------------+   +--------------+   +---------------+   +--------------+   +--------------+
|               |   |         |    |   |   |           |   |         |    |   |   |          |
|   (r1,c1)     |   |         |    |   |   |           |   |         |    |   |   |          |
|   +------+    |   |         |    |   |   |           |   +---------+    |   +---+          |
|   |      |    | = |         |    | - |   |           | - |      (r1,c2) | + |   (r1,c1)    |
|   |      |    |   |         |    |   |   |           |   |              |   |              |
|   +------+    |   +---------+    |   +---+           |   |              |   |              |
|        (r2,c2)|   |       (r2,c2)|   |   (r2,c1)     |   |              |   |              |
+---------------+   +--------------+   +---------------+   +--------------+   +--------------+

O(1) time
O(mn) space
*/
type NumMatrix struct {
	dp [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return NumMatrix{dp: [][]int{}}
	}
	dp := make([][]int, len(matrix)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(matrix[0])+1)
	}
	// dp[i+1][j+1] represents the sum of area from matrix[0][0] to matrix[i][j]
	// 不能用 dp[r][c]代表 因为一减会导致边被减去
	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[0]); c++ {
			dp[r+1][c+1] = dp[r+1][c] + dp[r][c+1] + matrix[r][c] - dp[r][c]
		}
	}
	return NumMatrix{dp: dp}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.dp[row2+1][col2+1] - this.dp[row1][col2+1] - this.dp[row2+1][col1] + this.dp[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */

/*
 上面的 row2+1 col2+1 不太好理解 也可以用 dp[r][c]代表 rc 的sum 和
 这个时候就要考虑边界情况了 各种 edge case 不太好处理
 各种+1的情况就好办了 因为初始化的时候 默认都是0 不会有 key 小于0的情况 而且没初始化的默认就是0
*/

type NumMatrix struct {
	dp [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return NumMatrix{dp: [][]int{}}
	}
	dp := make([][]int, len(matrix)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(matrix[0])+1)
	}
	dp[0][0] = matrix[0][0]
	if len(matrix[0]) > 1 {
		dp[0][1] = matrix[0][0] + matrix[0][1]
	} else {
		dp[0][1] = 0
	}
	if len(matrix) > 1 {
		dp[1][0] = matrix[0][0] + matrix[1][0]
	} else {
		dp[1][0] = 0
	}
	for r := 1; r < len(matrix); r++ {
		for c := 1; c < len(matrix[0]); c++ {
			dp[r][c] = dp[r-1][c] + dp[r][c-1] + matrix[r][c] - dp[r-1][c-1]
		}
	}
	return NumMatrix{dp: dp}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if row1 == 0 && col1 > 0 {
		return this.dp[row2][col2] - this.dp[0][col2] - this.dp[row2][col1-1] + this.dp[0][col1-1]
	}
	if row1 > 0 && col1 == 0 {
		return this.dp[row2][col2] - this.dp[row1-1][col2] - this.dp[row2][0] + this.dp[row1-1][0]
	}
	if row1 == 0 && col1 == 0 {
		return this.dp[row2][col2]
	}
	return this.dp[row2][col2] - this.dp[row1-1][col2] - this.dp[row2][col1-1] + this.dp[row1-1][col1-1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
