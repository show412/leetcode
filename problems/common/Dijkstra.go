import "math"

func networkDelayTime(times [][]int, N int, K int) int {
	// generate the graph
	graph := make(map[int][][]int)
	for _, edge := range times {
		graph[edge[0]] = append(graph[edge[0]], []int{edge[1], edge[2]})
	}
	// generate the distance, dist means S in Dijkstra
	dist := make(map[int]int)
	for i := 1; i <= N; i++ {
		dist[i] = math.MaxInt32
	}
	dist[K] = 0
	// seen act as U in Dijkstra
	seen := make(map[int]bool)
	// the Dijkstra loop solution
	// start to refine the U and S
	for {
		canNode := -1
		canDist := math.MaxInt32
		for i := 1; i <= N; i++ {
			if _, ok := seen[i]; !ok && dist[i] < canDist {
				canDist = dist[i]
				canNode = i
			}
		}
		if canNode < 0 {
			break
		}
		seen[canNode] = true
		if _, ok := graph[canNode]; ok {
			for _, info := range graph[canNode] {
				dist[info[0]] = min(dist[info[0]], dist[canNode]+info[1])
			}
		}
	}
	// start to find the shortest distance
	// notice it's to find the max time to some one node.package medium
	// it means that it will finish all the singal transfer
	res := 0
	for _, cand := range dist {
		// it means there is node impossible to arrive
		if cand == math.MaxInt32 {
			return -1
		}
		res = max(res, cand)
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
