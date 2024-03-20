/*
 * @Author: hongwei.sun
 * @Date: 2021-01-22 18:45:52
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-03-20 17:09:25
 * @Description: file content
 */
// https://leetcode.com/problems/remove-nth-node-from-end-of-list/
/*
Given a linked list, remove the n-th node from the end of list and return its head.

Example:

Given linked list: 1->2->3->4->5, and n = 2.

After removing the second node from the end, the linked list becomes 1->2->3->5.
Note:

Given n will always be valid.

Follow up:

Could you do this in one pass?
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 fast slow two pointer solution
 fast is ahead of slow pointer distance is n
 when fast is at the end and it's nil, slow is at the n postion
 */
 func removeNthFromEnd(head *ListNode, n int) *ListNode {
	/*
		for this case [1] 1 expected []
	*/
	if head.Next == nil && n == 1 {
		return nil
	}
	fast := head
	slow := head
	i := 0
	for i < n {
		/*
			[1,2]
			2
			expected [2]
		when fast move n position and out of this link list
		means we just remove head from end of list
		*/
		if fast.Next == nil {
			head = head.Next
			return head
		}
		fast = fast.Next
		i++
	}
	for fast != nil && fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	/*
		[1,2]
		1
		expected [1]
		the reason that there is no the logic:
		if slow.Next != nil && slow.Next.Next != nil
	*/
	slow.Next = slow.Next.Next
	return head
}
