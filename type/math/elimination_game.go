/*
 * @Author: hongwei.sun
 * @Date: 2021-01-22 18:45:51
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-10 17:30:08
 * @Description: file content
 */
// https://leetcode.com/problems/elimination-game/description/
/*
 There is a list of sorted integers from 1 to n. Starting from left to right,
 remove the first number and every other number afterward until you reach the end of the list.

# Repeat the previous step again, but this time from right to left,
remove the right most number and every other number from the remaining numbers.

# We keep repeating the steps again, alternating left to right and right to left, until a single number remains.

# Find the last number that remains starting with a list of length n.

# Example:

# Input:
# n = 9,
# 1 2 3 4 5 6 7 8 9
# 2 4 6 8
# 2 6
# 6

# Output:
# 6
*/
/*
本题的关在在于寻找数组变化的规律。首先，不论从左至右还是从右至左删除数字，
每次删除操作后数组中元素的个数都会减半，当数组只剩下一个数字时，该数字就是我们最终的答案。
换句话说，每当数组减半之后，我们需要观察剩余数组中首位元素的值，直到剩余数组长度为1时，这个首位元素的值即是返回结果。
因为最先从左侧开始删除，因此，原数组的首位元素（最左侧）一定是会被删除的。
当从右侧开始删除时，数组首元素（最左侧）是否会被删除，取决于当前数组长度，如果数组长度是偶数，首元素不会被删除，否则首元素会被删除。
比如1，2，3，4，5，从右侧删除时，5，3，1会被删除，而当个数为偶数时，比如1，2，3，4，这时我们只会删除4和2，首元素1不会被删除。
因此，数组首元素被删除的可能有两种，第一，从左侧开始删除。第二，从右侧开始删除并且数组长度为奇数。
当首元素被删除后，剩余数组的首元素就会变为它被删除前的下一个元素，那么问题来了，我们如何得知下一个元素是谁？
其实并不难，因为数组中的数字是存在规律的，初始时，数组中每两个数字之间间隔为1，进行一次删除后，间隔变为2，
再经过一次删除后间隔变为4，间隔会成倍增长。
只要我们知道当前经过了多少次删除操作，就能算出，两数之间的间隔，从而也就能知道被删除首元素的下一个元素的值。
*/
func lastRemaining(n int) int {
	leftToRight := true
	step := 1
	remain := n
	head := 1
	for remain > 1 {
		if leftToRight == true || remain%2 == 1 {
			head = head + step
		}
		step = step * 2
		remain = remain / 2
		leftToRight = !leftToRight
	}
	return head
}