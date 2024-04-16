/*
 * @Author: hongwei.sun
 * @Date: 2024-04-16 11:28:32
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-16 11:28:32
 * @Description: file content
 */
// https://leetcode.com/problems/remove-duplicates-from-sorted-list/
/*
Given a sorted linked list,
delete all duplicates such that each element appear only once.

Example 1:

Input: 1->1->2
Output: 1->2
Example 2:

Input: 1->1->2->3->3
Output: 1->2->3
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 这个题主要还是要理解head和dummy的指向 他们都指向一个内存地址
// 在开始的判断时，dummy和head都指向一个内存地址，当dummy的next变了 head的next也变了
 func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
		return head
	}
	dummy := head
	for dummy != nil && dummy.Next != nil {
		if dummy.Val == dummy.Next.Val {
			// 这步是关键
			dummy.Next = dummy.Next.Next
		} else {
			dummy = dummy.Next
		}
	}
	return head
}


func deleteDuplicates(head *ListNode) *ListNode {
	dumy := &ListNode{0, head}
	head = dumy
	for head.Next != nil && head.Next.Next != nil {
		if head.Next.Val == head.Next.Next.Val {
			val := head.Next.Val
			for head.Next != nil && head.Next.Val == val {
				head.Next = head.Next.Next
			}
		} else {
			head = head.Next
		}
	}
	return dumy.Next
}
