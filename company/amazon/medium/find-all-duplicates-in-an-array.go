// https://leetcode.com/problems/find-all-duplicates-in-an-array/
/*
Given an array of integers, 1 ≤ a[i] ≤ n (n = size of array),
some elements appear twice and others appear once.

Find all the elements that appear twice in this array.

Could you do it without extra space and in O(n) runtime?

Example:
Input:
[4,3,2,7,8,2,3,1]

Output:
[2,3]
*/
// TC O(n), SC O(1)
// 把数字放到改放的位置 如果它本身就在 move forward 如果不在就交换使其去正确的位置
// 但是不 move forward以免遗漏继续处理当前位置的数
// 下一次如果找到一个一样的数字 但是本身在该在的位置已经存在一个数了
// 说明这个是重复的 因为只找2个 这个时候把当前的数位置置位-1
func findDuplicates(nums []int) []int {
	res := make([]int, 0)
	if len(nums) == 0 {
		return res
	}
	i := 0
	for i < len(nums) {
		num := nums[i]
		if nums[i] == i+1 {
			i++
			continue
		}
		if num == -1 {
			i++
			continue
		}
		if nums[num-1] == num {
			nums[i] = -1
			res = append(res, num)
			i++
			continue
		}
		nums[num-1], nums[i] = nums[i], nums[num-1]
	}
	return res
}

// TC O(n), SC O(n)
func findDuplicates(nums []int) []int {
	array := make([]int, len(nums)+1)
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if array[num] == 1 {
			res = append(res, num)
		}
		array[num]++
	}
	return res
}

