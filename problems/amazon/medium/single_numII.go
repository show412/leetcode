// https://leetcode.com/problems/single-number-ii/
/*
Given a non-empty array of integers,
every element appears three times except for one, which appears exactly once.
Find that single one.

Note:

Your algorithm should have a linear runtime complexity.
Could you implement it without using extra memory?

Example 1:

Input: [2,2,3,2]
Output: 3
Example 2:

Input: [0,1,0,1,0,1,99]
Output: 99
*/
// refer to https://leetcode.com/problems/single-number-ii/solution/
// Go语言取反方式和C语言不同，Go语言不支持~符号
func singleNumber(nums []int) int {
	one := 0
	two := 0
	for i := 0; i < len(nums); i++ {
		one = ^two & (one ^ nums[i])
		two = ^one & (two ^ nums[i])
	}
	return one
}

/*
Intuition

Now let's discuss \mathcal{O}(1)O(1) space solution by using three bitwise operators

\sim x \qquad \textrm{that means} \qquad \textrm{bitwise NOT}∼xthat meansbitwise NOT

x \& y \qquad \textrm{that means} \qquad \textrm{bitwise AND}x&ythat meansbitwise AND

x \oplus y \qquad \textrm{that means} \qquad \textrm{bitwise XOR}x⊕ythat meansbitwise XOR

XOR

Let's start from XOR operator which could be used to detect the bit which appears odd number of times: 1, 3, 5, etc.

XOR of zero and a bit results in that bit

0 \oplus x = x0⊕x=x

XOR of two equal bits (even if they are zeros) results in a zero

x \oplus x = 0x⊕x=0

and so on and so forth, i.e. one could see the bit in a bitmask only if it appears odd number of times.

fig

That's already great, so one could detect the bit which appears once, and the bit which appears three times. The problem is to distinguish between these two situations.

AND and NOT

To separate number that appears once from a number that appears three times let's use two bitmasks instead of one: seen_once and seen_twice.

The idea is to

change seen_once only if seen_twice is unchanged

change seen_twice only if seen_once is unchanged

fig

This way bitmask seen_once will keep only the number which appears once and not the numbers which appear three times.
*/
