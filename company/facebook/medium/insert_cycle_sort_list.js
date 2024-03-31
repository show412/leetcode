// https://leetcode.com/problems/insert-into-a-cyclic-sorted-list/
/*
Given a node from a cyclic linked list which is sorted in ascending order, write a function to insert a value into the list such that it remains a cyclic sorted list. The given node can be a reference to any single node in the list, and may not be necessarily the smallest value in the cyclic list.

If there are multiple suitable places for insertion, you may choose any place to insert the new value. After the insertion, the cyclic list should remain sorted.

If the list is empty (i.e., given node is null), you should create a new single cyclic list and return the reference to that single node. Otherwise, you should return the original given node.

The following example may help you understand the problem better:





In the figure above, there is a cyclic sorted list of three elements. You are given a reference to the node with value 3, and we need to insert 2 into the list.





The new node should insert between node 1 and node 3. After the insertion, the list should look like this, and we should still return node 3.
*/
/**
 * // Definition for a Node.
 * function Node(val,next) {
 *    this.val = val;
 *    this.next = next;
 * };
 */
/**
 * @param {Node} head
 * @param {number} insertVal
 * @return {Node}
 */
/*
这道题让我们在循环有序的链表中插入结点，要求插入结点后，链表仍保持有序且循环。
题目中强调了corner case的情况，就是当链表为空时，
我们插入结点即要生成一个新的循环有序链表，那么我们可以先处理这个特殊情况，
比较简单，就是新建一个结点，然后将next指针指向自己即可。
好，下面来看给定的链表不为空的情况，最常见的情况就是要插入的结点值在两个有序结点值[a, b]之间，
那么只要满足 a <= insertVal <= b 即可。
由于是循环有序的链表，结点值不会一直上升，到某一个结点的时候，是最大值，但是下一个结点就是最小值了，
就是题目中的例子，结点4到结点1的时候，就是下降了。
那么在这个拐点的时候，插入值insertVal就会有两种特殊的情况，
其大于等于最大值，或者小于等于最小值，比如插入值是5，或者0的时候，
这两种情况都插入在结点4之后，可以放一起处理。而若其小于最大值，或者大于最小值，就是上面那种一般的情况，
不会在这里处理，所以我们只要检测如果属于上面的两种情况之一，就break掉循环，进行插入结点处理即可
*/
var insert = function (head, insertVal) {
  // 1. Edge Case. 就是新建一个结点，然后将next指针指向自己即可
  if (head === null) {
    return new Node(insertVal);
  }

  let prev = head, cur = prev.next;
  // 2. Loop ends where we can't find an insert point, this happens all values are the same.
  while (cur != head) {
    // 3. We are at the sorted Pivot, we can decide if our insertion point is the new Pivot
    // 由于是循环有序的链表，结点值不会一直上升，到某一个结点的时候，是最大值，但是下一个结点就是最小值了
    // 其大于等于最大值，或者小于等于最小值，对题里的例子比如插入值是5，或者0的时候，
    if (prev.val > cur.val) {
      if (insertVal >= prev.val || insertVal <= cur.val) break;
    }

    // 4. finally we found an insertion point where prev.val < insertVal < cur.val
    // 最常见的情况就是要插入的结点值在两个有序结点值[a, b]之间，那么只要满足 a <= insertVal <= b 即可。
    if (prev.val < insertVal && insertVal <= cur.val) break;
    prev = cur;
    cur = cur.next;
  }

  prev.next = new Node(insertVal, cur);
  return head;
};
