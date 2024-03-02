import "sort"

// https://leetcode.com/problems/non-overlapping-intervals/solution/
/*
Given a collection of intervals,
find the minimum number of intervals
you need to remove to make the rest of the intervals non-overlapping.



Example 1:

Input: [[1,2],[2,3],[3,4],[1,3]]
Output: 1
Explanation: [1,3] can be removed and the rest of intervals are non-overlapping.
Example 2:

Input: [[1,2],[1,2],[1,2]]
Output: 2
Explanation: You need to remove two [1,2]
to make the rest of intervals non-overlapping.
Example 3:

Input: [[1,2],[2,3]]
Output: 0
Explanation: You don't need to remove any of the intervals
since they're already non-overlapping.


Note:

You may assume the interval's end point is always bigger than its start point.
Intervals like [1,2] and [2,3] have borders "touching"
but they don't overlap each other.
*/
/*
这道题给了我们一堆区间，
让求需要至少移除多少个区间才能使剩下的区间没有重叠，
那么首先要给区间排序，根据每个区间的 start 来做升序排序，
然后开始要查找重叠区间，判断方法是看如果前一个区间的 end 大于后一个区间的 start，
那么一定是重复区间，此时结果 res 自增1，我们需要删除一个，那么此时究竟该删哪一个呢，
为了保证总体去掉的区间数最小，我们去掉那个 end 值较大的区间.
而在代码中，
我们并没有真正的删掉某一个区间，而是用一个变量 last 指向上一个需要比较的区间，
我们将 last 指向 end 值较小的那个区间；如果两个区间没有重叠，
那么此时 last 指向当前区间，继续进行下一次遍历
*/

type Interval [][]int

func (this Interval) Len() int {
	return len(this)
}

func (this Interval) Less(i, j int) bool {
	return this[i][0] < this[j][0]
}

func (this Interval) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func eraseOverlapIntervals(intervals [][]int) int {
	res := 0
	last := 0
	sort.Sort(Interval(intervals))
	for i := 1; i < len(intervals); i++ {
		// by this time, it has interval when next's start is smaller than last's end
		if intervals[i][0] < intervals[last][1] {
			res++
			// so which should be removed at this time?
			// we should remove the one which has bigger scope
			// so if next's end is smaller, last point to next(i) otherwise keep current last
			if intervals[i][1] < intervals[last][1] {
				last = i
			}
		} else {
			// no overlap between next's start and last's end, step next item
			last = i
		}
	}
	return res
}
