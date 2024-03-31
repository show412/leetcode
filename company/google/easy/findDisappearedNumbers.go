import "math"

// https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/
/*
Given an array of integers where 1 ≤ a[i] ≤ n (n = size of array),
some elements appear twice and others appear once.

Find all the elements of [1, n] inclusive that do not appear in this array.

Could you do it without extra space and in O(n) runtime?
You may assume the returned list does not count as extra space.

Example:

Input:
[4,3,2,7,8,2,3,1]

Output:
[5,6]
*/
// it's a very brillant idea
/*
The basic idea is that we iterate through the input array
and mark elements as negative using nums[nums[i] -1] = -nums[nums[i]-1].
In this way all the numbers that we have seen will be marked as negative.
In the second iteration, if a value is not marked as negative,
it implies we have never seen that index before,
so just add it to the return list.
*/
func findDisappearedNumbers(nums []int) []int {
	var res []int
	for i := 0; i < len(nums); i++ {
		val := int(math.Abs(float64(nums[i]))) - 1
		if nums[val] > 0 {
			nums[val] = -nums[val]
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			res = append(res, i+1)
		}
	}
	return res
}

// this solution is put every nums to be on the right position
// then to find out which psotion has no right number
func findDisappearedNumbers(nums []int) []int {
	nlen := len(nums)
	for i := 0; i < nlen; {
		if nums[i] != i+1 && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			continue
		}
		i++
	}
	var res []int
	for i := 0; i < nlen; i++ {
		if nums[i] != i+1 {
			res = append(res, i+1)
		}
	}
	return res
}
