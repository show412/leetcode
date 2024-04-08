/*
 * @Author: hongwei.sun
 * @Date: 2021-01-22 18:45:52
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-08 22:42:18
 * @Description: file content
 */
// https://leetcode.com/problems/hand-of-straights/
/*
Alice has some number of cards and she wants to rearrange the cards into groups so that each group is of size groupSize, and consists of groupSize consecutive cards.

Given an integer array hand where hand[i] is the value written on the ith card and an integer groupSize, return true if she can rearrange the cards, or false otherwise.

 

Example 1:

Input: hand = [1,2,3,6,2,3,4,7,8], groupSize = 3
Output: true
Explanation: Alice's hand can be rearranged as [1,2,3],[2,3,4],[6,7,8]
Example 2:

Input: hand = [1,2,3,4,5], groupSize = 4
Output: false
Explanation: Alice's hand can not be rearranged into groups of 4.

 

Constraints:

1 <= hand.length <= 104
0 <= hand[i] <= 109
1 <= groupSize <= hand.length
*/
/*
1， 不能整除groupSize肯定不能组成group
2， 排序然后通过map去处理
*/
func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}
	m := make(map[int]int)
	sort.Ints(hand)
	for _, card := range hand {
		m[card]++
	}
	for i := 0; i<len(hand);i++ {
		if m[hand[i]] > 0 {
			for j := 0; j < groupSize; j++ {
				if m[hand[i]+j] <= 0 {
					return false
				} else {
					m[hand[i]+j]--
				}
			}
		}
	}
	return true
}
