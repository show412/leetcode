import "math"

// https://leetcode.com/problems/network-delay-time/
/*
There are N network nodes, labelled 1 to N.

Given times, a list of travel times as directed edges times[i] = (u, v, w), where u is the source node, v is the target node, and w is the time it takes for a signal to travel from source to target.

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
		// determine which node will be to out from S in this loop
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

/*
Dijkstra Algorithm
refer to https://blog.csdn.net/heroacool/article/details/51014824

class Solution {
    Map<Integer, Integer> dist;
    public int networkDelayTime(int[][] times, int N, int K) {
        Map<Integer, List<int[]>> graph = new HashMap();
        for (int[] edge: times) {
            if (!graph.containsKey(edge[0]))
                graph.put(edge[0], new ArrayList<int[]>());
            graph.get(edge[0]).add(new int[]{edge[1], edge[2]});
        }
        dist = new HashMap();
        for (int node = 1; node <= N; ++node)
            dist.put(node, Integer.MAX_VALUE);

        dist.put(K, 0);
        boolean[] seen = new boolean[N+1];

        while (true) {
            int candNode = -1;
            int candDist = Integer.MAX_VALUE;
            for (int i = 1; i <= N; ++i) {
                if (!seen[i] && dist.get(i) < candDist) {
                    candDist = dist.get(i);
                    candNode = i;
                }
            }

            if (candNode < 0) break;
            seen[candNode] = true;
            if (graph.containsKey(candNode))
                for (int[] info: graph.get(candNode))
                    dist.put(info[0],
                             Math.min(dist.get(info[0]), dist.get(candNode) + info[1]));
        }

        int ans = 0;
        for (int cand: dist.values()) {
            if (cand == Integer.MAX_VALUE) return -1;
            ans = Math.max(ans, cand);
        }
        return ans;
    }
}
*/
