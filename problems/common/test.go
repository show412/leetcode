package main

import (
	"fmt"
	"math"
	"sort"
	// "math"
)

func main() {
	/*
			[[240,446],[745,948],[345,136],[341,68],[990,165],[165,580],[133,454],[113,527]]
		[[696,140],[95,393],[935,185],[767,205],[387,767],[214,960],[804,710],[956,307]]
	*/

	res := assignBikes([][]int{[]int{0, 0}, []int{1, 1}, []int{2, 0}}, [][]int{[]int{1, 0}, []int{2, 2}, []int{2, 1}})
	fmt.Println(res)
}

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
	// fmt.Println(distanceArray)
	// fmt.Println(workerMap)

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
