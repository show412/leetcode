/*
 * @Author: hongwei.sun
 * @Date: 2024-03-31 22:03:24
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-03-31 22:04:04
 * @Description: file content
 */
import "container/heap"

// https://leetcode.com/problems/merge-k-sorted-lists/
/*
Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.

Example:

Input:
[
  1->4->5,
  1->3->4,
  2->6
]
Output: 1->1->2->3->4->4->5->6
*/
// straightfoward is to merge one by one
// TC is O(nk)  SC is O(1)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	res := lists[0]
	for i := 1; i < len(lists); i++ {
		res = mergeLists(res, lists[i])
	}
	return res
 }

 func mergeLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			cur = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			cur =l2
			l2 = l2.Next
		}
	}
	for l1 != nil {
		cur.Next = l1
	}
	for l2 != nil {
		cur.Next = l2
	}
	return dummy.Next
 }


// ListHeap is type of heap of `ListNode`
type ListHeap []*ListNode

func (h ListHeap) Len() int           { return len(h) }
func (h ListHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h ListHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push the item to the heap
func (h *ListHeap) Push(x interface{}) { *h = append(*h, x.(*ListNode)) }

//Pop the item from heap
func (h *ListHeap) Pop() interface{} { x := (*h)[len(*h)-1]; *h = (*h)[0 : len(*h)-1]; return x }

func mergeKLists(lists []*ListNode) *ListNode {
	h := &ListHeap{}
	// init the heap, push the first node of every list onto the heep
	for _, l := range lists {
		if l != nil {
			heap.Push(h, l)
		}
	}

	List := &ListNode{}
	// Pop the minimum item from the heap and push the next value
	for curr := List; h.Len() > 0; curr = curr.Next {
		curr.Next = heap.Pop(h).(*ListNode)
		if curr.Next.Next != nil {
			heap.Push(h, curr.Next.Next)
		}
	}
	return List.Next
}
