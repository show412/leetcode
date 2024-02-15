/*
 * @Author: hongwei.sun
 * @Date: 2024-02-15 14:38:05
 * @LastEditors: your name
 * @LastEditTime: 2024-02-15 15:09:27
 * @Description: file content
 */
// https://leetcode.com/problems/merge-two-sorted-lists/
/*
Merge two sorted linked lists and return it as a new list.
The new list should be made by splicing together the nodes of the first two lists.

Example:

Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 需要定义一个dummy node to avoid edge case
	dummy := &ListNode{}
	// pre是往下移动的指针 because pre and dummy are both from &ListNode at begining,
	// so they point to same variable
	pre := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			pre.Next = l1
			l1 = l1.Next
		} else {
			pre.Next = l2
			l2 = l2.Next
		}
		pre = pre.Next
	}
	if l1 == nil {
		pre.Next = l2
	} else {
		pre.Next = l1
	}
	// return the begining of dummy next as first node of merged listNode
	return dummy.Next
}
