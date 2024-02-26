/*
 * @Author: hongwei.sun
 * @Date: 2024-02-26 09:53:40
 * @LastEditors: your name
 * @LastEditTime: 2024-02-26 10:44:39
 * @Description: file content
 */
//  https://leetcode.com/problems/number-of-islands/description/
// Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return the number of islands.

// An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

// Example 1:

// Input: grid = [
//   ["1","1","1","1","0"],
//   ["1","1","0","1","0"],
//   ["1","1","0","0","0"],
//   ["0","0","0","0","0"]
// ]
// Output: 1
// Example 2:

// Input: grid = [
//   ["1","1","0","0","0"],
//   ["1","1","0","0","0"],
//   ["0","0","1","0","0"],
//   ["0","0","0","1","1"]
// ]
// Output: 3

// Constraints:

// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 300
// grid[i][j] is '0' or '1'.

// typical BFS algo solution

func numIslands(grid [][]byte) int {
	ret := 0
	// define one 2d array, value in i and j means if or not it has been visited
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[i]))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' && visited[i][j] != true {
				ret++
				bfs(grid, visited, i, j)
			}
		}
	}
	return ret
}

func bfs(grid [][]byte, visited [][]bool, i int, j int) {
	// define one queue to traverse '1' adjacent to i and j
	q := make([][2]int, 0)
	q = append(q, [2]int{i, j})
	visited[i][j] = true
	offsets := [4][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for _, offset := range offsets {
			x := cur[0] + offset[0]
			y := cur[1] + offset[1]
			if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[x]) && visited[x][y] != true && grid[x][y] == '1' {
				// if there is one '1' x, y adjacent to i and j, continue to traverse x and y
				q = append(q, [2]int{x, y})
				visited[x][y] = true
			}
		}
	}
}