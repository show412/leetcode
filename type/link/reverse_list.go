/*
 * @Author: hongwei.sun
 * @Date: 2024-04-05 22:16:23
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-11 12:13:10
 * @Description: file content
 */
// https://leetcode.com/problems/reverse-linked-list/description/
/*
Reverse a singly linked list.

Example:

Input: 1->2->3->4->5->NULL
Output: 5->4->3->2->1->NULL
Follow up:

A linked list can be reversed either iteratively or recursively. Could you implement both?
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	// there should be one pointer type
	var pre *ListNode
	cur := head
	for cur != nil {
		nextNode := cur.Next
		cur.Next = pre
		pre = cur
		cur = nextNode
	}
	return pre
}
