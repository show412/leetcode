// 给定一个环形数组（最后一个元素的下一个元素是数组的第一个元素），为每个元素打印下一个更大的元素。 数字x的下一个更大的数是数组中下一个遍历顺序中出现的第一个更大的数字，这意味着您可以循环搜索以查找其下一个更大的数字。 如果它不存在，则为此数字输出-1。

// 样例
// 例1:

// 输入: [1,2,1]
// 输出: [2,-1,2]
// 解释：第一个1的下一个更大的数字是2;
// 数字2找不到下一个更大的数字;
// 第二个1的下一个更大的数字需要循环搜索，答案也是2。
// 例2:

// 输入: [1]
// 输出: [-1]
// 解释：
// 数字1找不到下一个更大的数字
// 注意事项
// 给定数组的长度不超过10000。
// https://www.lintcode.com/problem/next-greater-element-ii/description
/**
 * @param nums: an array
 * @return: the Next Greater Number for every element
 */
//  use stack and reverse travers to the array,
// the important is to add current element into the stack
func nextGreaterElements(nums []int) []int {
	// Write your code here
	if len(nums) == 0 || len(nums) == 1 {
		return nums
	}
	stack := make([]int, len(nums))
	result := make([]int, len(nums))
	length := len(nums)
	for i := length - 1; i >= 0; i-- {
		stack = append(stack, nums[i])
	}
	for i := length - 1; i >= 0; i-- {
		result[i] = -1

		for len(stack) != 0 && (stack[len(stack)-1] <= nums[i]) {
			stack = stack[:len(stack)-1]
		}
		if len(stack) != 0 {
			result[i] = stack[len(stack)-1]
		}
		// the important is to add current element into the stack for the following element
		stack = append(stack, nums[i])
	}
	return result
}