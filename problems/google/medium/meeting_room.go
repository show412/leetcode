import (
	"math"
	"sort"
)

// https://leetcode.com/problems/meeting-rooms-ii/

func minMeetingRooms(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	m := make(map[int]int, 0)
	arr := make([]int, 0)
	res := 0
	for i := 0; i < len(intervals); i++ {
		item := intervals[i]
		if _, ok := m[item[0]]; ok {
			m[item[0]] += 1
		} else {
			m[item[0]] = 1
		}
		if _, ok := m[item[1]]; ok {
			m[item[1]] -= 1
		} else {
			m[item[1]] = -1
		}
		arr = append(arr, item[0])
		arr = append(arr, item[1])
	}
	cnt := 0
	// fmt.Println(m)
	sort.Ints(arr)
	for k, v := range arr {
		if k != 0 && v == arr[k-1] {
			continue
		}
		cnt += m[v]
		res = int(math.Max(float64(res), float64(cnt)))
	}
	return res
}