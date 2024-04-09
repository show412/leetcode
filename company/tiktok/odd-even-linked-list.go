/*
 * @Author: hongwei.sun
 * @Date: 2024-04-09 16:54:21
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-09 17:31:11
 * @Description: file content
 */
// https://leetcode.com/problems/odd-even-linked-list/description/
/*
Given the head of a singly linked list, group all the nodes with odd indices together followed by the nodes with even indices, and return the reordered list.

The first node is considered odd, and the second node is even, and so on.

Note that the relative order inside both the even and odd groups should remain as it was in the input.

You must solve the problem in O(1) extra space complexity and O(n) time complexity.

 

Example 1:


Input: head = [1,2,3,4,5]
Output: [1,3,5,2,4]
Example 2:


Input: head = [2,1,3,5,6,4,7]
Output: [2,3,6,7,1,5,4]
 

Constraints:

The number of nodes in the linked list is in the range [0, 104].
-106 <= Node.val <= 106
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func oddEvenList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
		return head
	}
	// oddHead and odd both point to same address
	// so when begining odd Next is updated, oddHead is also updated at same time
	oddHead := &ListNode{}
	evenHead := &ListNode{}
	odd := oddHead
	even := evenHead
	count := 1
	for head != nil {
		if count % 2 == 1 {
			odd.Next = head
			odd = odd.Next
		} else {
			even.Next = head
			even = even.Next
		}
		head = head.Next
		count++
	}
	even.Next = nil
	odd.Next = evenHead.Next
	return oddHead.Next
 }