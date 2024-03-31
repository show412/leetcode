// https://leetcode.com/problems/single-number-iii
/*
Given an array of numbers nums,
in which exactly two elements appear only once
and all the other elements appear exactly twice.
Find the two elements that appear only once.

Example:

Input:  [1,2,1,3,2,5]
Output: [3,5]
Note:

The order of the result is not important.
So in the above example, [5, 3] is also correct.
Your algorithm should run in linear runtime complexity.
Could you implement it using only constant space complexity?
*/
// use hashmap
func singleNumber(nums []int) []int {
	m := make(map[int]int, 0)
	for _, num := range nums {
		m[num]++
	}
	res := make([]int, 0)
	for k, v := range m {
		if v == 1 {
			res = append(res, k)
		}
	}
	return res
}

// use bitwise
func singleNumber(nums []int) []int {
	xor := 0
	for _, num := range nums {
		xor ^= num
	}
	lastBit := xor - (xor & (xor - 1))
	group0 := 0
	group1 := 0
	for i := 0; i < len(nums); i++ {
		if lastBit&nums[i] == 0 {
			group0 ^= nums[i]
		} else {
			group1 ^= nums[i]
		}
	}
	res := make([]int, 2)
	res[0] = group0
	res[1] = group1
	return res
}

