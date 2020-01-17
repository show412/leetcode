import "math"

// https://leetcode.com/problems/walls-and-gates/
/*
You are given a m x n 2D grid initialized with these three possible values.

-1 - A wall or an obstacle.
0 - A gate.
INF - Infinity means an empty room. We use the value 2^31 - 1 = 2147483647 to represent INF
as you may assume that the distance to a gate is less than 2147483647.
Fill each empty room with the distance to its nearest gate.
If it is impossible to reach a gate, it should be filled with INF.

Example:

Given the 2D grid:

INF  -1  0  INF
INF INF INF  -1
INF  -1 INF  -1
  0  -1 INF INF
After running your function, the 2D grid should be:

  3  -1   0   1
  2   2   1  -1
  1  -1   2  -1
  0  -1   3   4
*/
// this is typical BFS for a matrix
// 典型的矩阵找最小距离的解法
func wallsAndGates(rooms [][]int) {
	if rooms == nil || len(rooms) == 0 {
		return
	}
	inf := math.MaxInt32
	path := []int{0, 1, 0, -1, 0}
	queue := [][2]int{}
	// 先找到gate
	for i := 0; i < len(rooms); i++ {
		for j := 0; j < len(rooms[0]); j++ {
			if rooms[i][j] == 0 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}
	// gate 去每一个 inf 因为是 bfs 肯定是最小距离
	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			for k := 0; k < 4; k++ {
				r := cur[0] + path[k]
				c := cur[1] + path[k+1]
				if r < 0 || c < 0 || r >= len(rooms) || c >= len(rooms[0]) || rooms[r][c] != inf {
					continue
				}
				rooms[r][c] = rooms[cur[0]][cur[1]] + 1
				queue = append(queue, [2]int{r, c})
			}
		}
	}
	return
}

