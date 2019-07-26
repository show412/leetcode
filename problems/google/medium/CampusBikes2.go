import (
	"math"
	"sort"
)

// https://leetcode.com/problems/campus-bikes-ii/
/*
On a campus represented as a 2D grid, there are N workers and M bikes, with N <= M. Each worker and bike is a 2D coordinate on this grid.

We assign one unique bike to each worker so that the sum of the Manhattan distances between each worker and their assigned bike is minimized.

The Manhattan distance between two points p1 and p2 is Manhattan(p1, p2) = |p1.x - p2.x| + |p1.y - p2.y|.

Return the minimum possible sum of Manhattan distances between each worker and their assigned bike.



Example 1:



Input: workers = [[0,0],[2,1]], bikes = [[1,2],[3,3]]
Output: 6
Explanation:
We assign bike 0 to worker 0, bike 1 to worker 1. The Manhattan distance of both assignments is 3, so the output is 6.
Example 2:



Input: workers = [[0,0],[1,1],[2,0]], bikes = [[1,0],[2,2],[2,1]]
Output: 4
Explanation:
We first assign bike 0 to worker 0, then assign bike 1 to worker 1 or worker 2, bike 2 to worker 2 or worker 1. Both assignments lead to sum of the Manhattan distances as 4.


Note:

0 <= workers[i][0], workers[i][1], bikes[i][0], bikes[i][1] < 1000
All worker and bike locations are distinct.
1 <= workers.length <= bikes.length <= 10
*/

func assignBikes(workers [][]int, bikes [][]int) int {
	res := math.MaxInt64
	cur := 0
	visit := make(map[int]bool, len(bikes))
	dfs(visit, workers, 0, bikes, cur, &res)
	return res
}

func dfs(visit map[int]bool, workers [][]int, start int, bikes [][]int, cur int, res *int) {
	if start >= len(workers) {
		*res = min(cur, *res)
		return
	}
	if cur > *res {
		return
	}
	for i := 0; i < len(bikes); i++ {
		if visit[i] == true {
			continue
		}
		visit[i] = true
		dfs(visit, workers, start+1, bikes, cur+dis(workers[start], bikes[i]), res)
		visit[i] = false
	}
}

func dis(worker []int, bike []int) int {
	return int(math.Abs(float64(worker[0]-bike[0]))) + int(math.Abs(float64(worker[1]-bike[1])))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// TBD
type DistanceSlice []Distance

func (d DistanceSlice) Len() int {
	return len(d)
}

func (d DistanceSlice) Less(i, j int) bool {
	return d[i].dis < d[j].dis
}

func (d DistanceSlice) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

type Distance struct {
	workerIndex int
	bikeIndex   int
	dis         int
}

func assignBikes(workers [][]int, bikes [][]int) int {
	res := 0
	distanceArray := make([]Distance, 0)
	workerMap := make(map[int]bool, len(workers))
	bikeMap := make(map[int]bool, len(bikes))
	for i := 0; i < len(workers); i++ {
		for j := 0; j < len(bikes); j++ {
			item := Distance{
				workerIndex: i,
				bikeIndex:   j,
				dis:         int(math.Abs(float64(workers[i][0]-bikes[j][0]))) + int(math.Abs(float64(workers[i][1]-bikes[j][1])))}
			distanceArray = append(distanceArray, item)
		}
	}
	sort.Sort(DistanceSlice(distanceArray))
	dfs(workerMap, bikeMap, distanceArray, &res)

}
func dfs(workerMap map[int]bool, bikeMap map[int]bool, distanceArray []Distance, res *int) {

	for i := 0; i < len(distanceArray); i++ {
		cur := distanceArray[i]
		v1, ok1 := workerMap[cur.workerIndex]
		v2, ok2 := bikeMap[cur.bikeIndex]
		if !ok1 && !ok2 && v1 != true && v2 != true && workerCount < len(workers) {
			res += cur.dis
			workerCount++
			workerMap[cur.workerIndex] = true
			bikeMap[cur.bikeIndex] = true
		}
	}
	return res
}
