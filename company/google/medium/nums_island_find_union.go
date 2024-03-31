// https://leetcode.com/problems/number-of-islands/
// use find union to resolve the questions
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	m := len(grid)
	n := len(grid[0])
	f := make([]int, (m+1)*(n+1)+1)
	total := 0
	// islandNum := 0
	// tag relation
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// fmt.Println(grid[i][j] == '1')
			if grid[i][j] == '1' {
				total++
				f[n*i+j+1] = n*i + j + 1 // tag the 1 coordinates
			}
		}
	}
	// unity and find the parent
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				// check the condition of unity, it's the adjcent of 1
				if i > 0 && grid[i-1][j] == '1' {
					unity(f[n*i+j+1], f[n*(i-1)+j+1], f)
				}
				if i < m-1 && grid[i+1][j] == '1' {
					unity(f[n*i+j+1], f[n*(i+1)+j+1], f)
				}
				if j > 0 && grid[i][j-1] == '1' {
					unity(f[n*i+j+1], f[n*i+j-1+1], f)
				}
				if j < n-1 && grid[i][j+1] == '1' {
					unity(f[n*i+j+1], f[n*i+j+1+1], f)
				}

			}
		}
	}
	// check the set number
	setCnt := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && f[n*i+j+1] == n*i+j+1 {
				setCnt++
			}
		}
	}
	return setCnt

}

// 典型的并查集的find union 的写法
// f[x]是指 x 的 parent
func unity(x int, y int, f []int) []int {
	x = find(x, f)
	y = find(y, f)
	if x != y {
		f[x] = y
	}
	return f
}

func find(x int, f []int) int {
	if f[x] == x {
		return x
	} else {
		// 路径压缩 使之后的查找效率更高
		f[x] = find(f[x], f)
	}
	return f[x]
}
