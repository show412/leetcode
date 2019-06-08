/*
Give a m * n board with a value of 0 or 1. At each step we can turn a 1 into 0 if there is another 1 in the same row or column. Return the max number of 1 we can turn into 0.

Example
Example 1:

Input：[[1,1,1,1],[1,1,0,1],[1,0,1,0]]
Output：8
Explanation:
In the board
1, 1, 1, 1
1, 1, 0, 1
1, 0, 1, 0
We can remove 1 from right to left, from bottom to top, until there is only one 1 at (0, 0). Totally 8 removed.
Example 2:

Input：[[1,0],[1,0]]
Output：1
Explanation:
In the board
1, 0
1, 0
We can only remove (0,0) or (1,0)
Notice
n and m do not exceed 5050
*/
/**
 * @param mp: the board
 * @return: the max number of points we can remove
 */
// it's union find
/* 并查集算法
这里采用union find，默认应该是说都是从左到右再从上到下遍历
可以这样理解：对于要消除的点，必要有一个依赖点，这个依赖点是同一行或同一列的。
那么这样就会形成一棵树（一个点可以有多个依赖点，不过一个就够了）。
union find的father就相当于是树根的点。
这样按照叶子到树根的方向消除，最后只剩树根(union find的father)不能被消除。
一个矩阵有多棵树，每棵树的树根不能被消除，所以是cnt - set.size()。
*/
func getAns(mp [][]int) int {
	// Write your code here.
	fa := make([]int, 5005)
	line := make([]int, 3000)
	column := make([]int, 3000)
	cnt := 0
	m := len(mp)
	n := len(mp[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mp[i][j] == 1 {
				fa[i*n+j+1] = i*n + j + 1
				// 某一个点总得有一个依赖点 某一行也得有个依赖点 某一列也得有个依赖点
				// 默认应该是说都是从左到右再从上到下遍历 所以这个依赖点就是行最左或者列最下
				line[i] = i*n + j + 1   //第i行的最右的点
				column[j] = i*n + j + 1 //第j列的最下的点
				cnt++
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mp[i][j] == 1 {
				// 默认应该是说都是从左到右再从上到下遍历
				unity(i*n+j+1, line[i], fa)   //这个点和i行的parent应该是一个
				unity(i*n+j+1, column[j], fa) //这个点和j列的parent应该是一个
			}
		}
	}
	// fmt.Println(fa[0:12])
	setSize = 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mp[i][j] == 1 && fa[i*n+j+1] == i*n+j+1 {
				setSize++
			}
		}
	}
	// 一个矩阵有多棵树，每棵树的树根不能被消除，所以是cnt - set.size()。
	cnt = cnt - setSize
	return cnt
}

func unity(x int, y int, fa []int) []int {
	x = find(x, fa)
	y = find(y, fa)
	// fmt.Println(fa[0:13])
	fa[x] = y
	return fa
}

func find(x int, fa []int) int {
	if fa[x] == x {
		return x
	} else {
		fa[x] = find(fa[x], fa)
	}
	return fa[x]
}
