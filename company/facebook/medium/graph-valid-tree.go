// https://leetcode.com/problems/graph-valid-tree/
/*
Given n nodes labeled from 0 to n-1 and a list of undirected edges (each edge is a pair of nodes),
write a function to check whether these edges make up a valid tree.

Example 1:

Input: n = 5, and edges = [[0,1], [0,2], [0,3], [1,4]]
Output: true
Example 2:

Input: n = 5, and edges = [[0,1], [1,2], [2,3], [1,3], [1,4]]
Output: false
Note: you can assume that no duplicate edges will appear in edges.
Since all edges are undirected, [0,1] is the same as [1,0]
and thus will not appear together in edges.
*/
/*
每个结点有零个或多个子结点；没有父结点的结点称为根结点；
每一个非根结点有且只有一个父结点；除了根结点外，每个子结点可以分为多个不相交的子树；
*/
func validTree(n int, edges [][]int) bool {
	fa := make([]int, n)
	// 首先每一个点的父设置为自己
	for i := 0; i < len(fa); i++ {
		fa[i] = i
	}
	// 根据边unity 各个点的父
	for _, edge := range edges {
		unity(edge[0], edge[1], fa)
	}
	/* 上面虽然整理好了父关系 还不全面 比如 [[0,1],[2,3],[1,2]]
	3的父是2 但是根据[1,2] 2的父还是1 1的父是0 这时候要重新整理3的父为0
	所以要重新find一遍
	*/
	for i := 0; i < len(fa); i++ {
		find(i, fa)
	}
	// 如果每个节点有不同的父 那肯定不是树 说明有孤出去的点
	last := fa[0]
	for _, v := range fa {
		if v != last {
			return false
		}
	}
	// 每个节点需要相连 所以边的总数应该是n-1
	return len(edges) == n-1
}

func unity(x int, y int, fa []int) {
	x = find(x, fa)
	y = find(y, fa)
	fa[y] = x
	return
}

func find(x int, fa []int) int {
	if fa[x] == x {
		return x
	} else {
		fa[x] = find(fa[x], fa)
	}
	return fa[x]
}

/*
5
[[0,1],[0,4],[1,4],[2,3]] false

5
[[0,1], [0,2], [0,3], [1,4]] true

5
[[0,1], [1,2], [2,3], [1,3], [1,4]] false

3
[[2,0],[2,1]] true

这个题不能用拓扑排序 因为是无向图 所以用拓扑会很麻烦 不知道起点和终点的方向
会导致很多 case cover不到 所以用 find union 比较合适
*/

func validTree(n int, edges [][]int) bool {
	treeEdges := make(map[int][]int, 0)
	nodes := make(map[int]bool, 0)
	indgree := make([]int, n)
	// check if or not there is a single node
	for _, edge := range edges {
		s := edge[0]
		e := edge[1]
		nodes[s] = true
		nodes[e] = true
		treeEdges[s] = append(treeEdges[s], e)
		indgree[e]++
	}
	if len(nodes) != n {
		return false
	}
	// topological sort
	queue := make([]int, 0)
	side := 0
	for i := 0; i < n; i++ {
		if indgree[i] == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		if _, ok := treeEdges[node]; ok {
			for _, eNode := range treeEdges[node] {
				if indgree[eNode] > 0 {
					indgree[eNode]--
					side++
				}
				if indgree[eNode] == 0 {
					queue = append(queue, eNode)

				}
			}
		}
	}
	return side == n-1
}