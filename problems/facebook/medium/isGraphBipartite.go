// https://leetcode.com/problems/is-graph-bipartite/solution/
/*
Given an undirected graph, return true if and only if it is bipartite.

Recall that a graph is bipartite if we can split
it's set of nodes into two independent subsets A and B such
that every edge in the graph has one node in A and another node in B.

The graph is given in the following form: graph[i] is a list of indexes j
for which the edge between nodes i and j exists.
Each node is an integer between 0 and graph.length - 1.
There are no self edges or parallel edges:
graph[i] does not contain i, and it doesn't contain any element twice.

Example 1:
Input: [[1,3], [0,2], [1,3], [0,2]]
Output: true
Explanation:
The graph looks like this:
0----1
|    |
|    |
3----2
We can divide the vertices into two groups: {0, 2} and {1, 3}.
Example 2:
Input: [[1,2,3], [0,2], [0,1,3], [0,2]]
Output: false
Explanation:
The graph looks like this:
0----1
| \  |
|  \ |
3----2
We cannot find a way to divide the set of nodes into two independent subsets.
*/
// solution: https://leetcode.com/problems/is-graph-bipartite/solution/
// 关键就是相邻的边不能有一样的颜色 然后枚举和遍历
// define data structure queue
type Queue struct {
	data []int
}

func (q *Queue) isEmpty() bool {
	return len(q.data) == 0
}

func (q *Queue) add(e int) {
	q.data = append(q.data, e)
}

func (q *Queue) poll() int {
	e := q.data[0]
	q.data = q.data[1:]
	return e
}

func isBipartite(graph [][]int) bool {
	var n = len(graph)
	var color = make([]int, n)

	for i := 0; i < n; i++ {
		if color[i] == 0 {
			res := bfs(i, graph, color)
			if !res {
				return false
			}
		}
	}
	return true
}

func bfs(start int, graph [][]int, color []int) bool {
	queue := &Queue{}

	queue.add(start)
	// random a acolor for this node
	color[start] = 1

	for !queue.isEmpty() {
		var u = queue.poll()
		var c = color[u]
		for _, v := range graph[u] {
			if color[v] == c {
				return false
			}
			if color[v] == 0 {
				// assign color
				if c == 1 {
					color[v] = 2
				} else {
					color[v] = 1
				}
				queue.add(v)
			}
		}
	}

	return true
}
