/*
 * @Author: hongwei.sun
 * @Date: 2024-04-15 16:54:35
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-15 17:51:24
 * @Description: file content
 */
//  https://leetcode.com/problems/reverse-linked-list-ii/description/
/*
Given the head of a singly linked list and two integers left and right where left <= right, reverse the nodes of the list from position left to position right, and return the reversed list.

 

Example 1:


Input: head = [1,2,3,4,5], left = 2, right = 4
Output: [1,4,3,2,5]
Example 2:

Input: head = [5], left = 1, right = 1
Output: [5]
 

Constraints:

The number of nodes in the list is n.
1 <= n <= 500
-500 <= Node.val <= 500
1 <= left <= right <= n
 

Follow up: Could you do it in one pass?
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/*
 1， 记住left 和 right 开始时候的node和前一个node
 2， 当在left和right中间时翻转，否则就是往下走
 3， left开始的之前的node 连上翻转的最后一个node(pre), 开始的node Next是 right的下一个node(cur) 
 */
 func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
    var start *ListNode
    var startPre *ListNode
    cur := head
    pre := dummy
	i := 1
	for i <= right {
        if i == left {
            startPre = pre
			start = cur
		}
        if i > left {
            next := cur.Next
            cur.Next = pre
			pre = cur
            cur = next
		} else {
            pre = cur
		    cur = cur.Next
        }
		i++
	}
	startPre.Next = pre
	start.Next = cur
	return dummy.Next
 }