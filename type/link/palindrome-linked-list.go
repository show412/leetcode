/*
 * @Author: hongwei.sun
 * @Date: 2024-04-11 11:10:45
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-04-11 12:14:56
 * @Description: file content
 */
// https://leetcode.com/problems/palindrome-linked-list/
/*
Given the head of a singly linked list, return true if it is a 
palindrome
 or false otherwise.

 

Example 1:


Input: head = [1,2,2,1]
Output: true
Example 2:


Input: head = [1,2]
Output: false
 

Constraints:

The number of nodes in the list is in the range [1, 105].
0 <= Node.val <= 9
 

Follow up: Could you do it in O(n) time and O(1) space?
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//  1, find middle node 
//  2, reverse second half link
//  3, compare
 func isPalindrome(head *ListNode) bool {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	
	// reverse second half link from slow
	// 定义 *ListNode 而不是&ListNode{} 否则会有一个实际存在的node
	var pre *ListNode
	for slow != nil {
		next := slow.Next
        slow.Next = pre
		pre = slow
		slow = next 
	}
	// compare pre and head
	for pre != nil && head != nil {
		if head.Val != pre.Val {
		    return false
        }
        head = head.Next
		pre = pre.Next
	}
	return true
 }