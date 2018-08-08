# Implement an iterator over a binary search tree (BST). Your iterator will be initialized with the root node of a BST.

# Calling next() will return the next smallest number in the BST.

# Note: next() and hasNext() should run in average O(1) time and uses O(h) memory, where h is the height of the tree.

# Credits:
# Special thanks to @ts for adding this problem and creating all test cases.
# Definition for a binary tree node.
class TreeNode
    attr_accessor :val, :left, :right
    def initialize(val)
        @val = val
        @left, @right = nil, nil
    end
end

class BSTIterator
    def initialize(root)
        @stack = []
        push_all_left(root)
    end

    def push_all_left(node)
        while node
            @stack << node
            node = node.left
        end
    end

    def has_next
        !@stack.empty?
    end

    def next
        node = @stack.pop
        push_all_left(node.right)
        node.val
    end
end

# Your BSTIterator will be called like this:
# i, v = BSTIterator.new(root), []
# while i.has_next()
#    v << i.next
# end
