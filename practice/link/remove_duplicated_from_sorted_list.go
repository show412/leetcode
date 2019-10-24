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
