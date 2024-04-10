/*
 * @Author: hongwei.sun
 * @Date: 2024-04-10 19:48:42
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-10 20:49:41
 * @Description: file content
 */
//  https://leetcode.com/problems/largest-number/description/
/*
Given a list of non-negative integers nums, arrange them such that they form the largest number and return it.

Since the result may be very large, so you need to return a string instead of an integer.

 

Example 1:

Input: nums = [10,2]
Output: "210"
Example 2:

Input: nums = [3,30,34,5,9]
Output: "9534330"
 

Constraints:

1 <= nums.length <= 100
0 <= nums[i] <= 109
*/
/*
两两数字转成字符串相加再转为数字比较就可以了
*/
func largestNumber(nums []int) string {
    sort.Slice(nums, func(i, j int)bool {
		str1 := strconv.Itoa(nums[i])
		str2 := strconv.Itoa(nums[j])
		num1, _ := strconv.Atoi(str1+str2)
		num2, _ := strconv.Atoi(str2+str1)
		if num1 > num2 {
			return true
		}
		return false
	})
    str := ""
    if nums[0] == 0 {
        return "0"
    }
    for _, v := range nums {
        str += strconv.Itoa(v)
    }
	return str
}