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
				// 注意这里是visiting
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

//https://leetcode.com/problems/course-schedule-ii/
/*
	这个题实际是用拓扑排序, 下面的方法是找出环 和这个题和没关系
*/
type Graph struct {
	nodes []int
	edges map[int][]int
}

func generateGraph(num int, edges [][]int) Graph {
	g := Graph{nodes: make([]int, 0), edges: make(map[int][]int, 0)}
	for i := 0; i < num; i++ {
		g.nodes = append(g.nodes, i)
	}
	for i := 0; i < len(edges); i++ {
		s := edges[i][0]
		e := edges[i][1]
		g.edges[s] = append(g.edges[s], e)
	}
	return g
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	g := generateGraph(numCourses, prerequisites)
	res := make([]int, 0)
	if findPrerequisites(g, numCourses, &res) == false {
		return []int{}
	} else {
		// 这是按顺序找出一个环
		for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
			res[i], res[j] = res[j], res[i]
		}
		return res
	}
}

func findPrerequisites(g Graph, num int, res *[]int) bool {
	visited := make(map[int]bool, 0)
	visiting := make(map[int]bool, 0)
	list := make([]int, 0)
	for i := 0; i < num; i++ {
		if len(g.edges[i]) <= 0 {
			continue
		}
		if dfs(g, i, visited, visiting, &list, res) == true {
			return false
		}
	}
	return true
}

func dfs(g Graph, node int, visited map[int]bool, visiting map[int]bool, list *[]int, res *[]int) bool {
	if visited[node] != true {
		visited[node] = true
		visiting[node] = true
		*list = append(*list, node)
		for j := 0; j < len(g.edges[node]); j++ {
			neighbour := g.edges[node][j]
			if visited[neighbour] != true && dfs(g, neighbour, visited, visiting, list, res) {
				return true
			} else if visiting[neighbour] == true {
				return true
			}
		}
	}
	visiting[node] = false
	if len(*list) == len(g.nodes) {
		cpy := make([]int, len(g.nodes))
		copy(cpy, *list)
		*res = cpy
		return false
	}

	return false
}
