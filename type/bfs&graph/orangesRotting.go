/*
 * @Author: hongwei.sun
 * @Date: 2021-01-22 18:45:51
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-06 20:20:23
 * @Description: file content
 */
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
/*
这题的关键是用bfs 然后用一个fresh标记新鲜的数量，这样有两个好处
1， for q > 0 的时候 带上fresh 可以避免最后一次都已经rottan的循环，不用最后return res-1
2， 也不用最后再双层循环一下找fresh去判断了 还能cover [[0]]这样的case
*/
func orangesRotting(grid [][]int) int {
	res := 0
	queue := make([][2]int, 0)
	fresh := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				fresh++
			}
			if grid[i][j] == 2 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}
	offsets := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
	for len(queue) != 0 && fresh > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			q := queue[0]
			queue = queue[1:]
			for _, offset := range offsets {
				x := q[0]+offset[0]
				y := q[1]+offset[1]
				if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) && grid[x][y] == 1 {
					grid[x][y] = 2
					fresh--
					queue = append(queue, [2]int{x, y})
				}
			}
		}
		res++
	}
	// for i := 0; i < len(grid); i++ {
	// 	for j := 0; j < len(grid[0]); j++ {
	// 		if grid[i][j] == 1 {
	// 			return -1
	// 		}
	// 	}
	// }
	if fresh == 0 {
		return res
	}
	return -1
}
