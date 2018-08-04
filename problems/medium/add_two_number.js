/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 https://leetcode.com/problems/add-two-numbers/description/
 */
function addTwoNumbers(l1, l2) {
  const before = new ListNode();
  // the tail next node will be before's next node at the same time
  let tail = before;
  let c = 0;

  while (l1 || l2 || c) {
    const v1 = l1 ? l1.val : 0;
    const v2 = l2 ? l2.val : 0;
    const v = v1+v2+c;
    // tail's next will be the before's next
    tail.next = new ListNode(v%10);
    // It's the key, the tail will be a new copy, not different from the origin one.
    // but in this time, the before's next has point to the new tail
    tail = tail.next;
    c = v >= 10 ? 1 : 0;
    l1 = l1&&l1.next;
    l2 = l2&&l2.next;
  }

  return before.next;
}
