// https://leetcode.com/problems/island-perimeter/
/*
You are given a map in form of a two-dimensional integer grid where 1 represents land and 0 represents water.

Grid cells are connected horizontally/vertically (not diagonally). The grid is completely surrounded by water, and there is exactly one island (i.e., one or more connected land cells).

The island doesn't have "lakes" (water inside that isn't connected to the water around the island). One cell is a square with side length 1. The grid is rectangular, width and height don't exceed 100. Determine the perimeter of the island.



Example:

Input:
[[0,1,0,0],
 [1,1,1,0],
 [0,1,0,0],
 [1,1,0,0]]

Output: 16

Explanation: The perimeter is the 16 yellow stripes in the image below:
*/
/*
求周长的问题
*/
func islandPerimeter(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				res += getLength(i, j, grid)
			}
		}
	}
	return res
}

// 每有一个相邻的1 就会有一个边不算在周长里
// 要考虑边界的情况 边界是算在周长里的
func getLength(i int, j int, grid [][]int) int {
	length := 4
	if i-1 >= 0 && grid[i-1][j] == 1 {
		length -= 1
	}
	if j+1 < len(grid[0]) && grid[i][j+1] == 1 {
		length -= 1
	}
	if i+1 < len(grid) && grid[i+1][j] == 1 {
		length -= 1
	}
	if j-1 >= 0 && grid[i][j-1] == 1 {
		length -= 1
	}
	return length
}
