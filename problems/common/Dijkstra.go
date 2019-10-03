import "math"

// https://leetcode.com/problems/network-delay-time/
/*
There are N network nodes, labelled 1 to N.

Given times, a list of travel times as directed edges times[i] = (u, v, w),
where u is the source node, v is the target node, and w is the time it takes for a signal to travel from source to target.

Now, we send a signal from a certain node K. How long will it take for all nodes to receive the signal? If it is impossible, return -1.



Example 1:



Input: times = [[2,1,1],[2,3,1],[3,4,1]], N = 4, K = 2
Output: 2


Note:

N will be in the range [1, 100].
K will be in the range [1, N].
The length of times will be in the range [1, 6000].
All edges times[i] = (u, v, w) will have 1 <= u, v <= N and 0 <= w <= 100.
*/
// Dijkstra solution
// refer to https://blog.csdn.net/heroacool/article/details/51014824
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
			// determine which node will be to out from S in this loop
			if _, ok := seen[i]; !ok && dist[i] < canDist {
				canDist = dist[i]
				canNode = i
			}
		}
		// 没找到有对应的最小的点的时候 break循环
		if canNode < 0 {
			break
		}
		seen[canNode] = true
		if _, ok := graph[canNode]; ok {
			for _, info := range graph[canNode] {
				/*
							  canNode ---- N1
						  /
						/
					N3
					N1 -> N3 = dist[info[0]]
					N1 -> canNode -> N3 = dist[canNode]+info[1]
				*/
				dist[info[0]] = min(dist[info[0]], dist[canNode]+info[1])
			}
		}
	}
	// start to find the shortest distance
	// notice it's to find the max time to some one node
	// it means that it will finish all the singal transfer
	// 因为它是说全遍历完 到最后一个节点 所以是拿最大
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
