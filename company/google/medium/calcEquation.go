// https://leetcode.com/problems/evaluate-division/
/*
Equations are given in the format A / B = k, where A and B are variables represented as strings, and k is a real number (floating point number). Given some queries, return the answers. If the answer does not exist, return -1.0.

Example:
Given a / b = 2.0, b / c = 3.0.
queries are: a / c = ?, b / a = ?, a / e = ?, a / a = ?, x / x = ? .
return [6.0, 0.5, -1.0, 1.0, -1.0 ].

The input is: vector<pair<string, string>> equations, vector<double>& values, vector<pair<string, string>> queries , where equations.size() == values.size(), and the values are positive. This represents the equations. Return vector<double>.

According to the example above:

equations = [ ["a", "b"], ["b", "c"] ],
values = [2.0, 3.0],
queries = [ ["a", "c"], ["b", "a"], ["a", "e"], ["a", "a"], ["x", "x"] ].


The input is always valid. You may assume that evaluating the queries will result in no division by zero and there is no contradiction.
*/
// res := calcEquation([][]string{[]string{"a", "b"}, []string{"b", "c"}}, []float64{2.0, 3.0}, [][]string{[]string{"a", "c"}, []string{"b", "a"}, []string{"a", "e"}, []string{"a", "a"}, []string{"x", "x"}})
// find union 是怎么看出来这是一个find union的解法呢？
// 其实就是一个转换成统一的变量的过程 就是一个find union的过程
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	m := make(map[string]string, 0)
	res := make([]float64, len(queries))
	dist := make(map[string]float64, 0)

	for i := 0; i < len(equations); i++ {
		unity(m, dist, values[i], equations[i][0], equations[i][1])
	}

	for i := 0; i < len(queries); i++ {
		_, ok1 := m[queries[i][0]]
		_, ok2 := m[queries[i][1]]
		if !ok1 || !ok2 {
			res[i] = -1.0
			continue
		}
		// 父类相同 说明 dist里的值是以同一个底数为标准的 比如 a => 6.0, c=>1.0 因为a和c有相同的parent：c
		// 所以这里dist他们都是以c为标准的倍数
		if find(m, dist, queries[i][0]) == find(m, dist, queries[i][1]) {
			res[i] = float64(dist[queries[i][0]] / dist[queries[i][1]])
			continue
		}
		res[i] = -1.0
	}
	return res
}

func add(m map[string]string, dist map[string]float64, s string) {
	_, ok := m[s]

	if !ok {
		m[s] = s
		dist[s] = 1.0
	}

	return
}

func unity(m map[string]string, dist map[string]float64, value float64, s1 string, s2 string) {
	add(m, dist, s1)
	add(m, dist, s2)
	// 这里找到s1的parent s2的parent， 比如 a=2.0*b b是a的parent
	p1 := find(m, dist, s1)
	p2 := find(m, dist, s2)
	m[p1] = p2
	// 注意这里p1的dist是通过s1和s2算出来的
	// 这里记得是p1和p2的关系 d(p1)/d(p2) = d(s1)/d(s2) = value
	dist[p1] = float64((dist[s2] * value) / dist[s1])
	return
}

func find(m map[string]string, dist map[string]float64, s string) string {
	v, ok := m[s]

	if !ok {
		m[s] = s
	}

	if s != v {
		// 递归找parent  a => 2.0 * b, b=> 3.0 * c, c => c
		// 递归找到a => 6.0 * c
		parent := find(m, dist, v)
		m[s] = parent
		dist[s] = dist[s] * dist[v]
	}
	// return的是m[s]
	return m[s]
}

// DFS with graph
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := make(map[string]map[string]float64)
	for i := 0; i < len(equations); i++ {
		first, second := equations[i][0], equations[i][1]
		if _, ok := graph[first]; !ok {
			graph[first] = make(map[string]float64)
		}
		if _, ok := graph[second]; !ok {
			graph[second] = make(map[string]float64)
		}

		graph[first][second] = values[i]
		graph[second][first] = 1.0 / values[i]
	}

	ans := make([]float64, len(queries))
	for i, query := range queries {
		ans[i] = helper(graph, query)
	}

	return ans
}

func helper(graph map[string]map[string]float64, query []string) float64 {
	first, second := query[0], query[1]
	if _, ok1 := graph[first]; !ok1 {
		return -1.0
	}
	if _, ok2 := graph[second]; !ok2 {
		return -1.0
	}

	if first == second {
		return 1.0
	}

	visited := make(map[string]bool)
	return dfs(visited, graph, first, second)
}

func dfs(visited map[string]bool, graph map[string]map[string]float64, current string, target string) float64 {
	if current == target {
		return 1.0
	}

	if _, ok := visited[current]; ok {
		return -1.0
	}

	visited[current] = true
	neighbors := graph[current]
	for neighbor, val := range neighbors {
		ret := dfs(visited, graph, neighbor, target)
		if ret == -1.0 {
			continue
		}

		return ret * val
	}

	return -1.0
}
