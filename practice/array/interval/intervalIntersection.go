// https://leetcode.com/problems/interval-list-intersections/
/*
Given two lists of closed intervals,
each list of intervals is pairwise disjoint and in sorted order.

Return the intersection of these two interval lists.

(Formally, a closed interval [a, b] (with a <= b)
denotes the set of real numbers x with a <= x <= b.
The intersection of two closed intervals is a set of real numbers
that is either empty, or can be represented as a closed interval.
For example, the intersection of [1, 3] and [2, 4] is [2, 3].)
Example 1:

Input: A = [[0,2],[5,10],[13,23],[24,25]], B = [[1,5],[8,12],[15,24],[25,26]]
Output: [[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]
Reminder: The inputs and the desired output are lists of Interval objects,
and not arrays or lists.


Note:

0 <= A.length < 1000
0 <= B.length < 1000
0 <= A[i].start, A[i].end, B[i].start, B[i].end < 10^9
NOTE: input types have been changed on April 15, 2019.
Please reset to default code definition to get new method signature.
*/
/*
A and B have these scenarios:

A	-----
B		-----

A	-----
B	 ---

A		----
B	----

A	-----
B	      -----
*/
func intervalIntersection(A [][]int, B [][]int) [][]int {
	i := 0
	j := 0
	m := len(A)
	n := len(B)
	res := make([][]int, 0)
	for i < m && j < n {
		low := max(A[i][0], B[j][0])
		high := min(A[i][1], B[j][1])
		if low <= high {
			res = append(res, []int{low, high})
		}
		if A[i][1] < B[j][1] {
			i++
		} else {
			j++
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
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
