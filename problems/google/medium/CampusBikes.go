import (
	"math"
	"sort"
)

// https://leetcode.com/problems/campus-bikes/
/*
On a campus represented as a 2D grid, there are N workers and M bikes, with N <= M. Each worker and bike is a 2D coordinate on this grid.

Our goal is to assign a bike to each worker. Among the available bikes and workers, we choose the (worker, bike) pair with the shortest Manhattan distance between each other, and assign the bike to that worker. (If there are multiple (worker, bike) pairs with the same shortest Manhattan distance, we choose the pair with the smallest worker index; if there are multiple ways to do that, we choose the pair with the smallest bike index). We repeat this process until there are no available workers.

The Manhattan distance between two points p1 and p2 is Manhattan(p1, p2) = |p1.x - p2.x| + |p1.y - p2.y|.

Return a vector ans of length N, where ans[i] is the index (0-indexed) of the bike that the i-th worker is assigned to.



Example 1:



Input: workers = [[0,0],[2,1]], bikes = [[1,2],[3,3]]
Output: [1,0]
Explanation:
Worker 1 grabs Bike 0 as they are closest (without ties), and Worker 0 is assigned Bike 1. So the output is [1, 0].
Example 2:



Input: workers = [[0,0],[1,1],[2,0]], bikes = [[1,0],[2,2],[2,1]]
Output: [0,2,1]
Explanation:
Worker 0 grabs Bike 0 at first. Worker 1 and Worker 2 share the same distance to Bike 2, thus Worker 1 is assigned to Bike 2, and Worker 2 will take Bike 1. So the output is [0,2,1].


Note:

0 <= workers[i][j], bikes[i][j] < 1000
All worker and bike locations are distinct.
1 <= workers.length <= bikes.length <= 1000

Sort the elements by distance. In case of a tie, sort them by the index of the worker. After that, if there are still ties, sort them by the index of the bike.

Can you do this in less than O(nlogn) time, where n is the total number of pairs between workers and bikes?

*/

/*
use case:
		[[240,446],[745,948],[345,136],[341,68],[990,165],[165,580],[133,454],[113,527]]
		[[696,140],[95,393],[935,185],[767,205],[387,767],[214,960],[804,710],[956,307]]
*/

type DistanceSlice []Distance

func (d DistanceSlice) Len() int {
	return len(d)
}

func (d DistanceSlice) Less(i, j int) bool {
	mdi := int(math.Abs(float64(d[i].workerCoordinate[0]-d[i].bikeCoordinate[0]))) +
		int(math.Abs(float64(d[i].workerCoordinate[1]-d[i].bikeCoordinate[1])))
	mdj := int(math.Abs(float64(d[j].workerCoordinate[0]-d[j].bikeCoordinate[0]))) +
		int(math.Abs(float64(d[j].workerCoordinate[1]-d[j].bikeCoordinate[1])))
	if mdi == mdj && d[i].workerIndex == d[j].workerIndex {
		return d[i].bikeIndex < d[j].bikeIndex
	} else if mdi == mdj {
		return d[i].workerIndex < d[j].workerIndex
	} else {
		return mdi < mdj
	}
}

func (d DistanceSlice) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

type Distance struct {
	workerIndex      int
	bikeIndex        int
	workerCoordinate []int
	bikeCoordinate   []int
}

func assignBikes(workers [][]int, bikes [][]int) []int {
	res := make([]int, len(workers))
	distanceArray := make([]Distance, 0)
	workerMap := make(map[int]bool, len(workers))
	bikeMap := make(map[int]bool, len(bikes))
	for i := 0; i < len(workers); i++ {
		for j := 0; j < len(bikes); j++ {
			item := Distance{
				workerIndex:      i,
				bikeIndex:        j,
				workerCoordinate: workers[i],
				bikeCoordinate:   bikes[j]}
			distanceArray = append(distanceArray, item)
		}
	}
	sort.Sort(DistanceSlice(distanceArray))

	for i := 0; i < len(distanceArray); i++ {
		cur := distanceArray[i]
		v1, ok1 := workerMap[cur.workerIndex]
		v2, ok2 := bikeMap[cur.bikeIndex]
		if !ok1 && !ok2 && v1 != true && v2 != true {
			res[cur.workerIndex] = cur.bikeIndex
			workerMap[cur.workerIndex] = true
			bikeMap[cur.bikeIndex] = true
		}
	}

	return res
}
