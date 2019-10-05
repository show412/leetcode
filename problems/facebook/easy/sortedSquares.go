// https://leetcode.com/problems/squares-of-a-sorted-array/
/*
Given an array of integers A sorted in non-decreasing order,
return an array of the squares of each number, also in sorted non-decreasing order.



Example 1:

Input: [-4,-1,0,3,10]
Output: [0,1,9,16,100]
Example 2:

Input: [-7,-3,2,3,11]
Output: [4,9,9,49,121]


Note:

1 <= A.length <= 10000
-10000 <= A[i] <= 10000
A is sorted in non-decreasing order.
*/
func sortedSquares(A []int) []int {
	res := make([]int, 0)
	if A[len(A)-1] <= 0 {
		for i := len(A) - 1; i >= 0; i-- {
			res = append(res, A[i]*A[i])
		}
		return res
	}
	if A[0] >= 0 {
		for i := 0; i < len(A); i++ {
			res = append(res, A[i]*A[i])
		}
		return res
	}
	p1 := 0
	p2 := 0
	for i := 0; i < len(A); i++ {
		if A[i] >= 0 && A[i-1] <= 0 {
			p1 = i - 1
			p2 = i
			break
		}
	}
	for p1 >= 0 && p2 <= len(A)-1 {
		if A[p1]*A[p1] < A[p2]*A[p2] {
			res = append(res, A[p1]*A[p1])
			p1--
		} else {
			res = append(res, A[p2]*A[p2])
			p2++
		}
	}
	for p1 >= 0 {
		res = append(res, A[p1]*A[p1])
		p1--

	}
	for p2 <= len(A)-1 {
		res = append(res, A[p2]*A[p2])
		p2++

	}
	return res
}
