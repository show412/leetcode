// https://leetcode.com/problems/course-schedule/
/*
There are a total of n courses you have to take, labeled from 0 to n-1.

Some courses may have prerequisites,
for example to take course 0 you have to first take course 1,
which is expressed as a pair: [0,1]

Given the total number of courses and a list of prerequisite pairs,
is it possible for you to finish all courses?

Example 1:

Input: 2, [[1,0]]
Output: true
Explanation: There are a total of 2 courses to take.
             To take course 1 you should have finished course 0. So it is possible.
Example 2:

Input: 2, [[1,0],[0,1]]
Output: false
Explanation: There are a total of 2 courses to take.
						 To take course 1 you should have finished course 0,
						 and to take course 0 you should
             also have finished course 1. So it is impossible.
Note:

The input prerequisites is a graph represented by a list of edges,
not adjacency matrices. Read more about how a graph is represented.
You may assume that there are no duplicate edges in the input prerequisites.
*/
/*
这个就是解决有向图有没有环的问题, solution: https://leetcode.com/problems/course-schedule/discuss/58612/Standard-algorithm-of-detecting-cycle-in-directed-graph-Java-runtime-6m

golang图如何表示
https://blog.csdn.net/Gerald_Jones/article/details/80158868
https://www.khanacademy.org/computing/computer-science/algorithms/graph-representation/a/representing-graphs
*/

// 拓扑排序的方法
func canFinish(numCourses int, prerequisites [][]int) bool {
	edges := make(map[int][]int, 0)
	indgree := make([]int, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		s := prerequisites[i][0]
		e := prerequisites[i][1]
		edges[e] = append(edges[e], s)
		indgree[s]++
	}
	queue := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if indgree[i] == 0 {
			queue = append(queue, i)
		}
	}
	i := 0
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		i++
		for _, edge := range edges[node] {
			if indgree[edge] > 0 {
				indgree[edge]--
			}
			if indgree[edge] == 0 {
				queue = append(queue, edge)
			}
		}
	}
	if i == numCourses {
		return true
	}
	return false
}

// DFS 方法
type Graph struct {
	nodes []int         // 节点集
	edges map[int][]int // 表示有向边
}

func generateGraph(num int, edges [][]int) Graph {
	g := Graph{nodes: make([]int, 0), edges: make(map[int][]int, 0)}
	// generate node
	for i := 0; i < num; i++ {
		node := i
		g.nodes = append(g.nodes, node)
	}
	// generate graph direct edges
	for i := 0; i < len(edges); i++ {
		sNode := edges[i][0]
		eNode := edges[i][1]
		g.edges[sNode] = append(g.edges[sNode], eNode)
	}
	return g
}

func checkCycle(g Graph, num int) bool {
	// 不能用BFS算法 因为可能起点不是某个有向边的起点，即起点可能都是有向边的终点,在map里不存在key
	// 所以用图的DFS算法
	/*
		即为一条道走到黑的搜索策略，行不通退回来换另外一条道再走到黑，依次直到搜索完成。
		其过程简要来说是对每一个可能的分支路径深入到不能再深入为止，而且每个节点只能访问一次。
	*/
	/*
		visited 表示遍历完了的点
		visiting 表示在这次遍历中正在遍历的点
	*/
	visited := make(map[int]bool, 0)
	visiting := make(map[int]bool, 0)
	for i := 0; i < num; i++ {
		if dfs(g, i, visited, visiting) == true {
			return false
		}
	}
	return true
}

/*
	即为一条道走到黑的搜索策略，行不通退回来换另外一条道再走到黑，依次直到搜索完成。
	其过程简要来说是对每一个可能的分支路径深入到不能再深入为止，而且每个节点只能访问一次。
	如果有环意味着在这一次遍历中肯定存在正在访问的点既visiting[i] == true的点
*/
func dfs(g Graph, i int, visited map[int]bool, visiting map[int]bool) bool {
	if visited[i] != true {
		visited[i] = true
		visiting[i] = true
		for j := 0; j < len(g.edges[i]); j++ {
			neighbour := g.edges[i][j]
			if visited[neighbour] != true && dfs(g, neighbour, visited, visiting) {
				return true
			} else if visiting[neighbour] == true {
				return true
			}
		}
	}
	visiting[i] = false
	return false
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	g := generateGraph(numCourses, prerequisites)
	// check cycle with DFS
	res := checkCycle(g, numCourses)
	return res
}

// 这是BFS遍历图的方法
/* 广度指的是从一个节点开始 发散性地遍历 周围节点。
从某个节点出发，访问它的所有邻接节点，
再从这些节点出发，访问它们未被访问过得邻接节点…直到所有节点访问完毕。
写法 https://segmentfault.com/a/1190000015358961?utm_source=tag-newest
*/
/*
func checkCycle(g Graph) bool {

	visited := make(map[int]bool, 0)
	q := make([]int, 0)
	q = append(q, g.nodes[0])
	visited[g.nodes[0]] = true
	for len(q) != 0 {
		node := q[0]
		q = q[1:]
		visited[node] = true
		for _, neighbour := range g.edges[node] {
			if visited[neighbour] == true {
				return false
			}
			q = append(q, neighbour)
		}
	}
	return true
}
*/
