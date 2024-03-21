/*
 * @Author: hongwei.sun
 * @Date: 2021-01-22 18:45:52
 * @LastEditors: hongwei.sun
 * @LastEditTime: 2024-03-21 17:55:18
 * @Description: file content
 */
// https://leetcode.com/problems/add-two-numbers/
/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// var res *ListNode  it's a nil because the res is only a address
	// cur := res
	res := &ListNode{}
	cur := res
	carry := 0
	for l1 != nil || l2 != nil {
		val := 0
		if l1 != nil {
			val = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}
		if carry == 1 {
			val += carry
		}
		if val >= 10 {
			carry = 1
			val = val % 10
		} else {
			carry = 0
		}
		// cur.Val = val
		// cur = cur.Next
		cur.Next = &ListNode{Val: val}
		cur = cur.Next
	}
	if carry == 1 {
		cur.Next = &ListNode{Val: carry}
		cur = cur.Next
	}
	return res.Next
}
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	// res is the result list beginning, one null pointer, 
	// cur is used to move forward
	// but return should be res.Next
	cur := res
	carry := 0
	for l1 != nil && l2 != nil {
		var1 := l1.Val
		var2 := l2.Val
		sum := var1 + var2 + carry
		if sum >= 10 {
			sum = sum%10
			carry = 1
		} else {
			carry = 0
		}
        
		// how to assign one value and move forward. 
		cur.Next = &ListNode{Val: sum}
		cur = cur.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		var1 := l1.Val
		sum := var1 + carry
		if sum >= 10 {
			sum = sum%10
			carry = 1
		} else {
			carry = 0
		}
		cur.Next = &ListNode{Val: sum}
		cur = cur.Next
		l1 = l1.Next
	}
	for l2 != nil {
		var2 := l2.Val
		sum := var2 + carry
		if sum >= 10 {
			sum = sum%10
			carry = 1
		} else {
			carry = 0
		}
		cur.Next = &ListNode{Val: sum}
		cur = cur.Next
		l2 = l2.Next
        
	}
    if carry == 1 {
		cur.Next = &ListNode{Val: carry}
		cur = cur.Next
	}
	return res.Next
}