# Reverse a singly linked list.

# Example:

# Input: 1->2->3->4->5->NULL
# Output: 5->4->3->2->1->NULL
# Follow up:

# A linked list can be reversed either iteratively or recursively. Could you implement both?
# https://leetcode.com/problems/reverse-linked-list/description/

# Definition for singly-linked list.
# class ListNode
#     attr_accessor :val, :next
#     def initialize(val)
#         @val = val
#         @next = nil
#     end
# end

# @param {ListNode} head
# @return {ListNode}
def reverse_list(head)
  pre = nil
  cur = head
  while(!cur.nil?)
    next_node = cur.next
    cur.next = pre
    pre = cur
    cur = next_node
  end
  pre
end
