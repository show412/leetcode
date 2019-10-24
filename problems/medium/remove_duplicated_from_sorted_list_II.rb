# https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/
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
def delete_duplicates(head)
  return head if head.nil? || head.next.nil?
  dummy = ListNode.new(0)
  dummy.next = head
  head = dummy
  while !head.next.nil? && !head.next.next.nil?
    if head.next.val == head.next.next.val
      val = head.next.val
      # 这段逻辑非常重要，while起到了定位有多少个重复node的作用，
      # next = next.next 起到了删除这些node的作用
      while !head.next.nil? && head.next.val == val
        head.next = head.next.next
      end
    else
      head = head.next
    end
  end
  dummy.next
end
