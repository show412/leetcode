/*
 * @Author: hongwei.sun
 * @Date: 2021-01-22 18:45:52
 * @LastEditors: your name
 * @LastEditTime: 2024-03-03 17:58:47
 * @Description: file content
 */
import (
	"math"
	"sort"
)

// https://leetcode.com/problems/merge-intervals/
/*
Given a collection of intervals, merge all overlapping intervals.

Example 1:

Input: [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
Example 2:

Input: [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.
NOTE: input types have been changed on April 15, 2019.
Please reset to default code definition to get new method signature.
*/
//  TC O(nlogn), SC O(1) or O(n)
func merge(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}
	res := make([][]int, 0)
	sort.Slice(intervals, func(i, j int) bool {
		before := intervals[i]
		end := intervals[j]
		if before[0] < end[0] { //按照开始时间进行排序
			return true
		}
		return false
	})
	for i := 0; i < len(intervals); i++ {
		item := intervals[i]
		if len(res) == 0 || res[len(res)-1][1] < item[0] {
			res = append(res, item)
		} else {
			res[len(res)-1][1] = max(res[len(res)-1][1], item[1])
		}
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// quick sort solution
func merge(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}
	res := make([][]int, 0)
	intervals = quickSort(intervals, 0, len(intervals)-1)

	for i := 0; i < len(intervals); i++ {
		item := intervals[i]
		if len(res) == 0 || res[len(res)-1][1] < item[0] {
			res = append(res, item)
		} else {
			res[len(res)-1][1] = int(math.Max(float64(res[len(res)-1][1]), float64(item[1])))
		}
	}

	return res
}

func quickSort(intervalArray [][]int, s int, e int) [][]int {
	if s >= e {
		return intervalArray
	}
	left := s
	right := e
	pivot := intervalArray[left]
	for s != e {
		// 一定是end先走
		for s < e && intervalArray[e][0] >= pivot[0] {
			e--
		}
		for s < e && intervalArray[s][0] <= pivot[0] {
			s++
		}
		if s < e {
			intervalArray[s], intervalArray[e] = intervalArray[e], intervalArray[s]
		}
	}
	intervalArray[left] = intervalArray[s]
	intervalArray[s] = pivot
	quickSort(intervalArray, left, s-1)
	quickSort(intervalArray, s+1, right)
	return intervalArray
}
