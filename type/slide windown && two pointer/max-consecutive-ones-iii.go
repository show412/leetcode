// https://leetcode.com/problems/max-consecutive-ones-iii/
/*
Given an array A of 0s and 1s, we may change up to K values from 0 to 1.

Return the length of the longest (contiguous) subarray that contains only 1s.



Example 1:

Input: A = [1,1,1,0,0,0,1,1,1,1,0], K = 2
Output: 6
Explanation:
[1,1,1,0,0,1,1,1,1,1,1]
Bolded numbers were flipped from 0 to 1.  The longest subarray is underlined.
Example 2:

Input: A = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], K = 3
Output: 10
Explanation:
[0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
Bolded numbers were flipped from 0 to 1.  The longest subarray is underlined.


Note:

1 <= A.length <= 20000
0 <= K <= A.length
A[i] is 0 or 1
*/
// two pointers and slide window
func longestOnes(A []int, K int) int {
	if K == len(A) {
		return len(A)
	}
	s, e := 0, 0
	res := 0
	n := 0
	for e < len(A) {
		if A[e] == 0 {
			n++
		}
		for n > K {
			if A[s] == 0 {
				n--
			}
			s++
		}
		e++
		res = max(res, e-s)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}