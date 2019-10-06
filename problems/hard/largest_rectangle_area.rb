# https://leetcode.com/problems/largest-rectangle-in-histogram/

# @param {Integer[]} heights
# @return {Integer}

# 考点：

# 单调栈
# 题解：
# 用九章算法强化班中讲过的单调栈(stack)。维护一个单调递增栈，逐个将元素 push 到栈里。push 进去之前先把 >= 自己的元素 pop 出来。
# 每次从栈中 pop 出一个数的时候，就找到了往左数比它小的第一个数（当前栈顶）和往右数比它小的第一个数（即将入栈的数），
# 从而可以计算出这两个数中间的部分宽度 * 被pop出的数，就是以这个被pop出来的数为最低的那个直方向两边展开的最大矩阵面积。
# 因为要计算两个数中间的宽度，因此放在 stack 里的是每个数的下标。

def largest_rectangle_area(heights)
  return 0 if(heights.nil? || heights.length == 0)
  i = 0
  max = 0
  stack = []
  while(i <= heights.length) do
    cur = i==heights.length ? -1 : heights[i]
    while(stack.length != 0 && cur <= heights[stack.last])
      height = heights[stack.pop]
      # if length is 0, it means those numbers between the pop number and current number
      # are all added into the wide, so the wide is i;
      # the wide should be i - stack.last -1 because the stack is incremental, perhaps
      # it has been pop several nums which is bigger than the pop item but we should add
      # these items into this pop items' wide
      wide = stack.length == 0 ? i : i - stack.last - 1
      max = [height * wide, max].max
    end
    stack.push(i)
    i += 1
  end
  return max
end
