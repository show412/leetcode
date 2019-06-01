import "math"

// Share
// You want to build a house on an empty land which reaches all buildings in the shortest amount of distance. You can only move up, down, left and right. You are given a 2D grid of values 0, 1 or 2, where:

// Each 0 marks an empty land which you can pass by freely.
// Each 1 marks a building which you cannot pass through.
// Each 2 marks an obstacle which you cannot pass through.
// Example:

// Input: [[1,0,2,0,1],[0,0,0,0,0],[0,0,1,0,0]]

// 1 - 0 - 2 - 0 - 1
// |   |   |   |   |
// 0 - 0 - 0 - 0 - 0
// |   |   |   |   |
// 0 - 0 - 1 - 0 - 0

// Output: 7

// Explanation: Given three buildings at (0,0), (0,4), (2,2), and an obstacle at (0,2),
//              the point (1,2) is an ideal empty land to build a house, as the total
//              travel distance of 3+3+1=7 is minimal. So return 7.
// Note:
// There will be at least one building. If it is not possible to build such house according to the above rules, return -1.
// https://leetcode.com/problems/shortest-distance-from-all-buildings/
// 这种解法是在没有2的情况下 有了2就只能BFS了
// O(n*m)
// 我们可以借鉴的是通过两个数组记下哪些row和col是1，
// 然后可以用另一个数组提前算好row方向的distance和col方向的distance
// 按着算法 同一个row上的row距离是0 同理同一个col上的col距离是0
func shortestDistance(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return -1
	}
	res := -1
	rowHouse := make([]int, len(grid))
	colHouse := make([]int, len(grid[0]))
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				// 注意这一定是++ 因为可能一个row上有多个1 这样到这个row的距离就会被加权
				rowHouse[i]++
				colHouse[j]++
			}
		}
	}

	// every point distance to these points which is 1 in level of row and col
	rowDistance := make([]int, len(grid))
	colDistance := make([]int, len(grid[0]))
	getDistance(len(grid), rowHouse, rowDistance)
	getDistance(len(grid[0]), colHouse, colDistance)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				d := rowDistance[i] + colDistance[j]
				if d >= 0 {
					if res == -1 || res > d {
						// fmt.Println(i)
						// fmt.Println(j)
						// fmt.Println(d)
						res = d
					}
				}
			}
		}
	}
	return res

}

// 这个有些巧妙 目的是算i点到其他所有点j的距离 但是利用乘以house[j]是否等于1
// 来使distance只存了到右house点的距离
func getDistance(size int, house []int, distance []int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			distance[i] += house[j] * int(math.Abs(float64(j-i)))
		}
	}
}

/*
有障碍物的情况下 数组里有2只能用BFS 因为不能通过x和y的相减和相乘来巧妙的计算了
  用BFS 二维数组如何使用BFS:
  上下左右利用这个数组 然后遍历这样的code去实现
  final int[] shift = new int[] {0, 1, 0, -1, 0};
  int nextRow = curr[0] + shift[k];
  int nextCol = curr[1] + shift[k + 1];
*/
public class Solution {
    public int shortestDistance(int[][] grid) {
        if (grid == null || grid[0].length == 0) return 0;
				// 上下左右利用这个数组 然后遍历这样的code去实现
				// int nextRow = curr[0] + shift[k];
        // int nextCol = curr[1] + shift[k + 1];
				final int[] shift = new int[] {0, 1, 0, -1, 0};

        int row  = grid.length, col = grid[0].length;
        int[][] distance = new int[row][col];
        int[][] reach = new int[row][col];
        int buildingNum = 0;

        for (int i = 0; i < row; i++) {
            for (int j =0; j < col; j++) {
                if (grid[i][j] == 1) {
                    buildingNum++;
                    Queue<int[]> myQueue = new LinkedList<int[]>();
                    myQueue.offer(new int[] {i,j});

                    boolean[][] isVisited = new boolean[row][col];
                    int level = 1;

                    while (!myQueue.isEmpty()) {
                        int qSize = myQueue.size();
                        for (int q = 0; q < qSize; q++) {
                            int[] curr = myQueue.poll();

                            for (int k = 0; k < 4; k++) {
                                int nextRow = curr[0] + shift[k];
                                int nextCol = curr[1] + shift[k + 1];

                                if (nextRow >= 0 && nextRow < row && nextCol >= 0 && nextCol < col
                                    && grid[nextRow][nextCol] == 0 && !isVisited[nextRow][nextCol]) {
                                        //The shortest distance from [nextRow][nextCol] to this building
                                        // is 'level'.
                                        distance[nextRow][nextCol] += level;
                                        reach[nextRow][nextCol]++;

                                        isVisited[nextRow][nextCol] = true;
                                        myQueue.offer(new int[] {nextRow, nextCol});
                                    }
                            }
                        }
                        level++;
                    }
                }
            }
        }
				/*
					distance[i][j] 不一定和reach[i][j]就是相等的
					distance是记距离的 reach是记有多少个1可以到这
					因为level是每次queue之后都会加1
				*/
        int shortest = Integer.MAX_VALUE;
        for (int i = 0; i < row; i++) {
            for (int j = 0; j < col; j++) {
                if (grid[i][j] == 0 && reach[i][j] == buildingNum) {
                    shortest = Math.min(shortest, distance[i][j]);
                }
            }
        }

        return shortest == Integer.MAX_VALUE ? -1 : shortest;


    }
}

// this is typical BFS for a matrix
func shortestDistance(grid [][]int) int {
	if grid == nil || len(grid[0]) == 0 {
		return -1
	}
	path := []int{0, 1, 0, -1, 0}
	// distance := [len(grid)][len(grid[0])]int{}
	distance := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		distance[i] = make([]int, len(grid[0]))
	}
	reach := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		reach[i] = make([]int, len(grid[0]))
	}
	buildingNum := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				buildingNum++
				queue := [][2]int{}
				queue = append(queue, [2]int{i, j})
				visit := make(map[[2]int]bool)
				level := 1
				for len(queue) > 0 {
					size := len(queue)
					for z := 0; z < size; z++ {
						cur := queue[len(queue)-1]
						queue = queue[:len(queue)-1]
						for k := 0; k < 4; k++ {
							nextRow := cur[0] + path[k]
							nextCol := cur[1] + path[k+1]
							if nextRow >= 0 && nextRow < len(grid) && nextCol >= 0 && nextCol < len(grid[0]) && grid[nextRow][nextCol] == 0 && visit[[2]int{nextRow, nextCol}] != true {
								visit[[2]int{nextRow, nextCol}] = true
								queue = append(queue, [2]int{nextRow, nextCol})
								distance[nextRow][nextCol] += level
								reach[nextRow][nextCol]++
							}
						}
					}
					level++
				}
			}
		}
	}
	const MAX = int(^uint(0) >> 1)
	res := MAX
	fmt.Println(buildingNum)
	fmt.Println(distance)
	fmt.Println(reach)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 && reach[i][j] == buildingNum {
				res = int(math.Min(float64(res), float64(distance[i][j])))
			}
		}
	}
	if res == MAX {
		return -1
	}
	return res
}