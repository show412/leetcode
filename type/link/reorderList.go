/*
 * @Author: hongwei.sun
 * @Date: 2021-01-22 18:45:51
 * @LastEditors: your name
 * @LastEditTime: 2024-03-07 23:21:30
 * @Description: file content
 */
// https://leetcode.com/problems/reorder-list/
/*
Given a singly linked list L: L0→L1→…→Ln-1→Ln,
reorder it to: L0→Ln→L1→Ln-1→L2→Ln-2→…

You may not modify the values in the list's nodes, only nodes itself may be changed.

Example 1:

Given 1->2->3->4, reorder it to 1->4->2->3.
Example 2:

Given 1->2->3->4->5, reorder it to 1->5->2->4->3.
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/*
1, find middle
2, reverse second half and divide two link
3, reorder two link one by one
*/
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	// find middle
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// find second half link
	secondHalfHead := slow.Next
	// divide two link by set first Half next is nil
	slow.Next = nil
	// reverse second half
	var pre *ListNode
	for secondHalfHead != nil {
		secondHalfNext := secondHalfHead.Next
		secondHalfHead.Next = pre
		pre = secondHalfHead
		secondHalfHead = secondHalfNext
	}
	// merge first half and reversed second half one by one
	firstHalf := head
	secondHalf := pre
	var next1 *ListNode
	var next2 *ListNode
	for firstHalf != nil && secondHalf != nil {
		next1 = firstHalf.Next
		next2 = secondHalf.Next
		firstHalf.Next = secondHalf
		secondHalf.Next = next1
		firstHalf = next1
		secondHalf = next2
	}
	return
}

/*
	这样不行 因为是指针传递 全部反转之后 head实际已经指向最后一个node
	实际只有一个link 原来的link并没有记下来
  取得反转的一个link 取得link的长度l.
	正常和反转的link轮流取一个node 直到反转的link取到l/2长
	好像不是解释的这样
*/
// func reorderList(head *ListNode) {
// 	l := 0
// 	rev := head
// 	var pre *ListNode
// 	for rev != nil {
// 		l++
// 		next := rev.Next
// 		rev.Next = pre
// 		pre = rev
// 		rev = next
// 	}
// 	revHead := pre
// 	n := 0
// 	var next1 *ListNode
// 	var next2 *ListNode
// 	cur := head
// 	for n <= l/2 {
// 		next1 = cur.Next
// 		next2 = revHead.Next
// 		cur.Next = revHead
// 		revHead.Next = next1
// 		cur = next1
// 		revHead = next2
// 		n++
// 	}
// 	return
// }
