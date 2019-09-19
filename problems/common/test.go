package main

import (
	"fmt"
)

/*
test case:
4
[[2,0],[1,0],[3,1],[3,2],[1,3]]
false
*/
func main() {
	res := canFinish(4, [][]int{[]int{2, 0}, []int{1, 0}, []int{3, 1}, []int{3, 2}, []int{1, 3}})
	// res := canFinish(4, [][]int{[]int{0, 1}})
	fmt.Println(res)
}

type Graph struct {
	nodes []int         // 节点集
	edges map[int][]int // 邻接表表示的无向图
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
	// 不能用BFS算法 因为可能起点不是一个边的起点，所以用图的DFS算法
	/*
		即为一条道走到黑的搜索策略，行不通退回来换另外一条道再走到黑，依次直到搜索完成。
		其过程简要来说是对每一个可能的分支路径深入到不能再深入为止，而且每个节点只能访问一次。
	*/
	visited := make(map[int]bool, 0)
	visiting := make(map[int]bool, 0)
	for i := 0; i < num; i++ {
		// if visited[i] != true && len(g.edges[i]) > 0 {
		if dfs(g, i, visited, visiting) == true {
			return false
		}
		// }
	}
	return true
}

func dfs(g Graph, i int, visited map[int]bool, visiting map[int]bool) bool {
	if visited[i] != true {
		visited[i] = true
		visiting[i] = true
		for j := 0; j < len(g.edges[i]); j++ {
			neighbour := g.edges[i][j]
			if visited[neighbour] != true && dfs(g, neighbour, visited, visiting) {
				fmt.Println("1111")
				fmt.Println(visited)
				fmt.Println(neighbour)
				return true
			} else if visiting[neighbour] == true {
				fmt.Println("2222")
				fmt.Println(visiting)
				fmt.Println(neighbour)
				return true
			}
		}
	}
	visiting[i] = false
	return false
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	g := generateGraph(numCourses, prerequisites)
	// check cycle
	// visited := make(map[int]bool, numCourses)
	fmt.Println(g)
	res := checkCycle(g, numCourses)
	return res
}

// 这是BFS遍历图的方法
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
