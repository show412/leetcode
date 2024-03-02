/*
 * @Author: hongwei.sun
 * @Date: 2024-03-03 00:25:41
 * @LastEditors: your name
 * @LastEditTime: 2024-03-03 00:25:42
 * @Description: file content
 */
/*
 * @Author: hongwei.sun
 * @Date: 2024-03-02 23:46:11
 * @LastEditors: your name
 * @LastEditTime: 2024-03-03 00:23:48
 * @Description: file content
 */
//  https://leetcode.com/problems/longest-consecutive-sequence/
// Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

// You must write an algorithm that runs in O(n) time.

// Example 1:

// Input: nums = [100,4,200,1,3,2]
// Output: 4
// Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
// Example 2:

// Input: nums = [0,3,7,2,5,8,4,6,0,1]
// Output: 9

// Constraints:

// 0 <= nums.length <= 105
// -109 <= nums[i] <= 109
/*
thought is divide different consecutive sequence and find the longest
how to divide different consecutive sequence:
	1, using hash to check if i of nums has left neighbor and right neighbor
	2, if it has no left neighbor, it's as one start of sequence
	3, then we go to find right neighbor in this hash map,
	4, then to compare these different consecutive sequence length to find max LCS
*/
func longestConsecutive(nums []int) int {
	longest := 0
	m := make(map[int]bool, len(nums))
	for _, num := range nums {
		m[num] = true
	}
	/*
		this is to iterate nums, we can iterate in array, and also can iterate in map
		in array, first is index, second is value; in map first is key, second is value of key of map
		iterate array cost more time, iterate map is more efficent
	*/
	for _, num := range nums {
		if _, ok := m[num-1]; !ok {
			l := 0
			// check if exist num+l in nums by checking m
			for m[num+l] == true {
				l++
			}
			longest = max(longest, l)
		}
	}
	return longest
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}