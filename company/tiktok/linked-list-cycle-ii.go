/*
 * @Author: hongwei.sun
 * @Date: 2024-04-09 23:33:40
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-19 11:30:47
 * @Description: file content
 */
// https://leetcode.com/problems/linked-list-cycle-ii/description/
 /*
 Given the head of a linked list, return the node where the cycle begins. If there is no cycle, return null.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to (0-indexed). It is -1 if there is no cycle. Note that pos is not passed as a parameter.

Do not modify the linked list.

 

Example 1:


Input: head = [3,2,0,-4], pos = 1
Output: tail connects to node index 1
Explanation: There is a cycle in the linked list, where tail connects to the second node.
Example 2:


Input: head = [1,2], pos = 0
Output: tail connects to node index 0
Explanation: There is a cycle in the linked list, where tail connects to the first node.
Example 3:


Input: head = [1], pos = -1
Output: no cycle
Explanation: There is no cycle in the linked list.
 

Constraints:

The number of the nodes in the list is in the range [0, 104].
-105 <= Node.val <= 105
pos is -1 or a valid index in the linked-list.
 */
 /*
 这个是要找出这个环的起点
 fast 指针⼀次都 2 步，slow 指针⼀次⾛ 1 步。令链表 head 到环的第⼀个点
需要 x1 步，从环的第⼀个点到相遇点需要 x2 步，从环中相遇点回到环的第⼀个点需要 x3 步。那么环
的总⻓度是 x2 + x3 步。
fast 和 slow 会相遇，说明他们⾛的时间是相同的，可以知道他们⾛的路程有以下的关系：
fast 的 t = (x1 + x2 + x3 + x2) / 2
slow 的 t = (x1 + x2) / 1
x1 + x2 + x3 + x2 = 2 * (x1 + x2)
所以 x1 = x3
所以 2 个指针相遇以后，如果 slow 继续往前⾛， 指针回到起点 head，两者都每次⾛⼀步，那么必
定会在环的起点相遇，相遇以后输出这个点即是结果。
 */

 /**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	isCycle := false
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		// this is key, if there is one cycle, fast can't move out from list
		fast = fast.Next.Next
		if slow == fast {
			isCycle = true
			break
		}
	}
	if isCycle == false {
		return nil
	}
	for head != slow {
		head = head.Next
		slow = slow.Next
	}
	return slow
}