/*
 * @Author: hongwei.sun
 * @Date: 2021-01-22 18:45:51
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-10 11:22:30
 * @Description: file content
 */
// https://leetcode.com/problems/majority-element-ii/
/*
Given an integer array of size n, find all elements that appear more than ⌊ n/3 ⌋ times.

Note: The algorithm should run in linear time and in O(1) space.

Example 1:

Input: [3,2,3]
Output: [3]
Example 2:

Input: [1,1,1,3,3,2,2,2]
Output: [1,2]
*/
// TC O(n), SC O(n)
func majorityElement(nums []int) []int {
	result := make([]int, 0)
	m := make(map[int]int)

	for _, val := range nums {
		if v, ok := m[val]; ok {
			m[val] = v + 1
		} else {
			m[val] = 1
		}
	}

	for key, value := range m {
		if value > len(nums)/3 {
			result = append(result, key)
		}
	}
	return result
}

// TC O(n), SC O(1)
// Boyer–Moore_majority_vote_algorithm 一种变体 
// https://zh.wikipedia.org/wiki/%E5%A4%9A%E6%95%B0%E6%8A%95%E7%A5%A8%E7%AE%97%E6%B3%95
// 只所以要进行第二次遍历比较，是因为只能说找到了两个数是众数，但是不一定是大于1/3
func majorityElement(nums []int) []int {
	// Since we are checking if a num appears more than 1/3 of the time
	// it is only possible to have at most 2 nums (>1/3 + >1/3 = >2/3)
	count1, count2, candidate1, candidate2 := 0, 0, 0, 1

	// Select candidates
	for _, num := range nums {
		if num == candidate1 {
			count1++
		} else if num == candidate2 {
			count2++
		} else if count1 <= 0 {
			// We have a bad first candidate, replace!
			candidate1, count1 = num, 1
		} else if count2 <= 0 {
			// We have a bad second candidate, replace!
			candidate2, count2 = num, 1
		} else {
			// Both candidates suck, boo!
			count1--
			count2--
		}
	}

	// Recount! Since in the first loop the count was used as
	// a weight to know if we should choose a different candidate
	// it may not be the total count. We must recount for the
	// final 2 candidates
	count1 = 0
	count2 = 0
	for _, num := range nums {
		if num == candidate1 {
			count1++
		} else if num == candidate2 {
			count2++
		}
	}

	length := len(nums)
	if count1 > length/3 && count2 > length/3 {
		return []int{candidate1, candidate2}
	}

	if count1 > length/3 {
		return []int{candidate1}
	}

	if count2 > length/3 {
		return []int{candidate2}
	}

	return []int{}
}