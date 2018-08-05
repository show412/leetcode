// Given a linked list, remove the n-th node from the end of list and return its head.

// Example:

// Given linked list: 1->2->3->4->5, and n = 2.

// After removing the second node from the end, the linked list becomes 1->2->3->5.
// Note:

// Given n will always be valid.

// Follow up:

// Could you do this in one pass?
// https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/
/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} head
 * @param {number} n
 * @return {ListNode}
 */
 // Time complexity : O(L).
 // The algorithm makes two traversal of the list, first to calculate list length L and second to find the (L - n) th node.
 // There are 2L-n operations and time complexity is O(L).
 // Space complexity : O(1).
var removeNthFromEnd = function(head, n) {
    let tail = head;
    let length = 1;
    if(head.next == null && n==1){
      return null;
    }
   while(tail.next !== null){
    tail = tail.next;
    length ++;
   }
   tail = head;
   if(length == n){
     return head.next;
   }else{
     for(i=1; i<length - n ; i++){
      tail = tail.next;
     }
     tail.next = tail.next.next;
     return head;
   }
};

// In one pass
let left = head;
let right = head;
let before;

while(n--) {

    // remove head - n is equal to size of linked list
    if(right.next === null) {
        head = head.next;
        return head;
    }

    right = right.next;

  // n exceeded the size of linked list
  if(!right) return;
}

while(right) {
  before = left;
  right = right.next;
  left = left.next;
}
before.next = before.next.next;
return head;
}
