/*
 * @Author: hongwei.sun
 * @Date: 2024-04-05 22:18:19
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-05 23:12:52
 * @Description: file content
 */
//  https://leetcode.com/problems/insert-interval/description/
/*
You are given an array of non-overlapping intervals intervals where intervals[i] = [starti, endi] represent the start and the end of the ith interval and intervals is sorted in ascending order by starti. You are also given an interval newInterval = [start, end] that represents the start and end of another interval.

Insert newInterval into intervals such that intervals is still sorted in ascending order by starti and intervals still does not have any overlapping intervals (merge overlapping intervals if necessary).

Return intervals after the insertion.

Note that you don't need to modify intervals in-place. You can make a new array and return it.

 

Example 1:

Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
Output: [[1,5],[6,9]]
Example 2:

Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
Output: [[1,2],[3,10],[12,16]]
Explanation: Because the new interval [4,8] overlaps with [3,5],[6,7],[8,10].
 

Constraints:

0 <= intervals.length <= 104
intervals[i].length == 2
0 <= starti <= endi <= 105
intervals is sorted by starti in ascending order.
newInterval.length == 2
0 <= start <= end <= 105
*/
/*
两个数组interval 有以下5种可能
1,        -----
   --
2,       ----- 
      ---- 
3,       ------
          ---
4,      -----
           ----
5,      -----
              ----		   

但实际就分为三种：  1) 在前面 2）overlap 3）在后面
在前面的话 直接insert 然后后面的都insert进来
overlap的话 通过 [min, max] 取的一个新的newinterval 不用append 然后根据循环去后面再按一样的逻辑去处理
在后面的话， 当前的insert进来， 但是不知道会不会和后面的有overlap 然后根据循环去后面再按一样的逻辑去处理
关键是要有考虑后面的思想 不能拘泥于当前一个元素上
*/

func insert(intervals [][]int, newInterval []int) [][]int {
    res := make([][]int, 0)
	for i := 0; i < len(intervals); i++ {
		interval := intervals[i]
		if newInterval[1] < interval[0] {
			res = append(res, newInterval)
			// ...是展开后面的参数
			res = append(res, intervals[i:]...)
			return res
		} else if newInterval[0] > interval[1] {
			// newInterval可能还是后面的overlap 不用append 
			res = append(res, interval)
		} else {
			newInterval = []int{min(newInterval[0], interval[0]), max(newInterval[1], interval[1])}
		}
	}
	res = append(res, newInterval)
	return res
}

func max(a, b int) int {
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
