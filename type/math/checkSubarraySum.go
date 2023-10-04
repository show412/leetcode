// https://leetcode.com/problems/continuous-subarray-sum/
/*
Given a list of non-negative numbers and a target integer k,
write a function to check if the array has a continuous subarray of size at least 2
that sums up to a multiple of k, that is, sums up to n*k where n is also an integer.

Example 1:

Input: [23, 2, 4, 6, 7],  k=6
Output: True
Explanation: Because [2, 4] is a continuous subarray of size 2 and sums up to 6.
Example 2:

Input: [23, 2, 6, 4, 7],  k=6
Output: True
Explanation: Because [23, 2, 6, 4, 7] is an continuous subarray of size 5 and sums up to 42.


Note:

The length of the array won't exceed 10,000.
You may assume the sum of all the numbers is in the range of a signed 32-bit integer.
*/
/*
也可以用前缀和的方式解决 但是效率低
Time complexity : O(n^2)O(n). Two for loops are used for considering every subarray possible.
Space complexity : O(n)O(n). sumsum array of size nn is used.

the theory is that:
a%k = x
b%k = x
(a - b) %k = x -x = 0
here a - b = the sum between i and j.
[23,2,3,1,5,3]
6

index ---- value ------sum ----- modulo % k
0----------23----------23---------- 5
1----------2-----------25---------- 1
2----------3-----------28---------- 4
3----------1-----------29---------- 5
matches with index = 0 whose value is 23. Thus, its possbile

Time complexity : O(n). Only one traversal of the array numsnums is done.
Space complexity : O(min(n,k)). The HashMap can contain upto min(n,k)min(n,k) different pairings.
*/
// common question
func checkSubarraySum(nums []int, k int) bool {
	modMap := make(map[int]int, 0)
	sumArray := make([]int, len(nums))
	sumArray[0] = nums[0]
	// 当 k == 0的时候, modMap 存和 因为这时候如果后面有一些和是相等的 说明中间这些的连续和是0
	if k != 0 {
		modMap[nums[0]%k] = 0
	} else {
		modMap[nums[0]] = 0
	}
	for i := 1; i < len(nums); i++ {
		sumArray[i] = nums[i] + sumArray[i-1]
		// 从 index 0 开始就符合的情况
		if sumArray[i] == k || (k != 0 && sumArray[i]%k == 0) {
			return true
		}
		mod := sumArray[i]
		if k != 0 {
			mod = sumArray[i] % k
		}
		if v, ok := modMap[mod]; ok {
			if i-v > 1 {
				return true
			} else {
				// 这种情况就是 num[i]是0的情况 这时候是不能认为 true 的
				// 因为题目要求at least 2 numbers
				// [0,1,0] 0
				continue
			}
		}
		modMap[mod] = i
	}
	return false
}

