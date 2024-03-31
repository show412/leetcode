// https://leetcode.com/problems/rotting-oranges/
/*
In a given grid, each cell can have one of three values:

the value 0 representing an empty cell;
the value 1 representing a fresh orange;
the value 2 representing a rotten orange.
Every minute, any fresh orange that is adjacent (4-directionally) to a rotten orange
becomes rotten.

Return the minimum number of minutes that must elapse
until no cell has a fresh orange.  If this is impossible, return -1 instead.



Example 1:



Input: [[2,1,1],[1,1,0],[0,1,1]]
Output: 4
Example 2:

Input: [[2,1,1],[0,1,1],[1,0,1]]
Output: -1
Explanation:  The orange in the bottom left corner (row 2, column 0) is never rotten, because rotting only happens 4-directionally.
Example 3:

Input: [[0,2]]
Output: 0
Explanation:  Since there are already no fresh oranges at minute 0, the answer is just 0.


Note:

1 <= grid.length <= 10
1 <= grid[0].length <= 10
grid[i][j] is only 0, 1, or 2.
*/
// it's a matrix BFS problem
func orangesRotting(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	if len(grid) == 1 && len(grid[0]) == 1 {
		if grid[0][0] == 1 {
			return -1
		}
		return 0
	}
	res := 0
	// first find out the rotten orange posion and push them into a queue
	queue := make([][2]int, 0)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}
	// BFS to find the fresh orange
	path := []int{0, 1, 0, -1, 0}
	for len(queue) != 0 {
		size := len(queue)
		res++
		// 一定要有这个 因为这算一步
		for i := 0; i < size; i++ {
			r := queue[0][0]
			c := queue[0][1]
			queue = queue[1:]
			for k := 0; k < 4; k++ {
				r1 := r + path[k]
				c1 := c + path[k+1]
				if r1 >= 0 && c1 >= 0 && r1 < len(grid) && c1 < len(grid[0]) && grid[r1][c1] == 1 {
					grid[r1][c1] = 2
					queue = append(queue, [2]int{r1, c1})
				}
			}
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}
	return res - 1
}
