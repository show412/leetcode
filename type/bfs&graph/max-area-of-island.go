/*
 * @Author: hongwei.sun
 * @Date: 2024-03-22 16:50:30
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-03-22 16:50:30
 * @Description: file content
 */
/*
 * @Author: hongwei.sun
 * @Date: 2021-01-22 18:45:52
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-03-22 16:49:54
 * @Description: file content
 */
// https://leetcode.com/problems/max-area-of-island/
/*
Given a non-empty 2D array grid of 0's and 1's, an island is a group of 1's (representing land) connected 4-directionally (horizontal or vertical.) You may assume all four edges of the grid are surrounded by water.

Find the maximum area of an island in the given 2D array. (If there is no island, the maximum area is 0.)

Example 1:

[[0,0,1,0,0,0,0,1,0,0,0,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,1,1,0,1,0,0,0,0,0,0,0,0],
 [0,1,0,0,1,1,0,0,1,0,1,0,0],
 [0,1,0,0,1,1,0,0,1,1,1,0,0],
 [0,0,0,0,0,0,0,0,0,0,1,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,0,0,0,0,0,0,1,1,0,0,0,0]]
Given the above grid, return 6. Note the answer is not 11, because the island must be connected 4-directionally.
Example 2:

[[0,0,0,0,0,0,0,0]]
Given the above grid, return 0.
Note: The length of each dimension in the given grid does not exceed 50.
*/
func maxAreaOfIsland(grid [][]int) int {
	res := 0
	// this is to record if position is visitted, it's a 2 dimension array
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[i]))
	}
    for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			// need to check if it is 1 and don't visitted it, 
			// we go to bfs, if it's 0, skip it
            if grid[i][j] == 1 && visited[i][j] == false {
                res = max(bfs(i,j,visited,grid)+1, res)
            }
			
		}
	}
	return res
}
func bfs(i int, j int, visited [][]bool, grid [][]int) int {
	q := make([][2]int,0)
	q = append(q, [2]int{i,j})
	visited[i][j] = true
	// notice define one array, it's to use {} even 2d array
	offset := [4][2]int{{1,0},{0,1},{-1,0},{0,-1}}
	res := 0
	for len(q) != 0 {
		item := q[0]
		q = q[1:]
		// traverse one 2d array with range, first param is index in for loop, second is value.
		for _, direction := range offset {
			x := direction[0] + item[0]
			y := direction[1] + item[1]
			if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) && grid[x][y] == 1 && visited[x][y] == false {
				res++
				visited[x][y] = true
				q = append(q, [2]int{x,y})
			}
		}
	}
	return res
}

func max (a, b int) int {
	if a > b {
		return a
	}
	return b
}