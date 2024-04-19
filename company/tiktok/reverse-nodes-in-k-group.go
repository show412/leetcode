/*
 * @Author: hongwei.sun
 * @Date: 2024-04-19 11:08:16
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-19 11:08:16
 * @Description: file content
 */
//  https://leetcode.com/problems/reverse-nodes-in-k-group/description/
// Given the head of a linked list, reverse the nodes of the list k at a time, and return the modified list.

// k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes, in the end, should remain as it is.

// You may not alter the values in the list's nodes, only nodes themselves may be changed.

 

// Example 1:


// Input: head = [1,2,3,4,5], k = 2
// Output: [2,1,4,3,5]
// Example 2:


// Input: head = [1,2,3,4,5], k = 3
// Output: [3,2,1,4,5]
 

// Constraints:

// The number of nodes in the list is n.
// 1 <= k <= n <= 5000
// 0 <= Node.val <= 1000
/*
1，each k nodes to be one group, and reverse group by group
2, we need to one function to find kth node
3, one pointer for previous node of beginning node in group. one pointer for next node of last node in group
difficulty is on how to define and orgnaize these nodes
4， reverse的时候 pre的初始值是下一个group的beginning
5， after group reverseing, need to put groupPre to be end of this group for nex group use. 
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reverseKGroup(head *ListNode, k int) *ListNode {
    dummy := &ListNode{0, head}
	groupPre := dummy
	for {
		// because it's one loop, groupPre is updating, so we use groupPre here as param
		// kth is from groupPre to find(pre head), not from head
		kth := getKth(groupPre, k)
		if kth == nil {
			break
		}
		// one pointer for next node of last node in this group
		groupNext := kth.Next
		// reverse node in group
		// initialize pre be groupNext to make firt cur to point to next beginning of nex group
		pre, cur := groupNext, groupPre.Next
		for cur != groupNext {
			tmp := cur.Next
			cur.Next = pre
			pre = cur
			cur = tmp
		}
		/* 
		this is most difficult part in this problem, 
		我们已经把这个group都reverse了，我们需要把groupPre 指向这个group的最后一个，为下一个group使用
		这里是把groupPre 赋值到reverse之后的group里的最后一个node
		*/
		tmp := groupPre.Next
		groupPre.Next = kth
		groupPre = tmp
	}
	// because all variable is pointer, AKA memory address
	/*
	用dummy.next是因为指向的地址是不变的，这块地址的值已经被reverse换成了别的node
	head是一个变量 值是当时的第一个node的内容 一系列操作后仍然是这个值 比如一直是1 所以这里不能返回head
	*/
	return dummy.Next
 }

 func getKth(node *ListNode, k int) *ListNode {
	for k > 0 && node != nil {
		node = node.Next
		k--
	}
	return node
 }