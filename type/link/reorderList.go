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
//  1, find middle 2, reverse second half and divide two link 3, reorder two link one by one
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	// find middle
	fHalf := head
	sHalf := head
	// 注意这里都是第二部分的Next
	for sHalf.Next != nil && sHalf.Next.Next != nil {
		fHalf = fHalf.Next
		sHalf = sHalf.Next.Next
	}
	// reverse second half
	cur2 := fHalf.Next
	// divide two link by set first Half next is nil
	fHalf.Next = nil
	var pre2 *ListNode
	for cur2 != nil {
		next2 := cur2.Next
		cur2.Next = pre2
		pre2 = cur2
		cur2 = next2
	}
	// record one by one
	fHalf = head
	sHalf = pre2
	var next1 *ListNode
	var next2 *ListNode
	for fHalf != nil && sHalf != nil {
		next1 = fHalf.Next
		next2 = sHalf.Next
		fHalf.Next = sHalf
		sHalf.Next = next1
		fHalf = next1
		sHalf = next2
	}
	return
}

/*
	这样不行 因为是指针传递 全部反转之后 head实际已经指向最后一个node
	实际只有一个link 原来的link并没有记下来
  取得反转的一个link 取得link的长度l.
	正常和反转的link轮流取一个node 直到反转的link取到l/2长
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
