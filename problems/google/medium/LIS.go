import (
	"math"
	"sort"
)

// https://leetcode.com/problems/longest-increasing-subsequence/
// Given an unsorted array of integers, find the length of longest increasing subsequence.

// Example:

// Input: [10,9,2,5,3,7,101,18]
// Output: 4
// Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
// Note:

// There may be more than one LIS combination, it is only necessary for you to return the length.
// Your algorithm should run in O(n2) complexity.
// Follow up: Could you improve it to O(n log n) time complexity?
// DP is O(n^2) and SC is O(n)
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 1
	}
	f := make([]int, len(nums))
	f[0] = 1
	res := 1
	for i := 1; i < len(nums); i++ {
		f[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				f[i] = int(math.Max(float64(f[i]), float64(f[j]+1)))
			}
		}
		res = int(math.Max(float64(f[i]), float64(res)))
	}
	return res
}

// use the binary search to make the TC O(nlogn)
/* It means we replace the corresponding number with the smaller number and
add the bigger number at the bottome of dp array to keep on the following elements
could be checked.
It's a hard to understand solution.
*/
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	const MAX = int(^uint(0) >> 1)
	f := make([]int, len(nums))
	// init the every element is MAX in f array as dp
	for i := 1; i < len(f); i++ {
		f[i] = MAX
	}
	f[0] = nums[0]
	// binary search: find the smallest element which is bigger or equal to nums[i]
	// the index is j and replace the j as nums[i]
	for i := 1; i < len(nums); i++ {
		j := sort.Search(len(f), func(j int) bool { return f[j] >= nums[i] })
		f[j] = nums[i]
	}
	// the f length which is not MAX is the result
	for i := len(f) - 1; i >= 0; i-- {
		if f[i] != MAX {
			return i + 1
		}
	}
	return 0
}
