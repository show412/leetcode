import "container/heap"

// https://leetcode.com/problems/merge-k-sorted-lists/
/*
Merge k sorted linked lists and return it as one sorted list.
Analyze and describe its complexity.

Example:

Input:
[
  1->4->5,
  1->3->4,
  2->6
]
Output: 1->1->2->3->4->4->5->6
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// ListHeap is type of heap of `ListNode`
// it seems the priority queue solution
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
