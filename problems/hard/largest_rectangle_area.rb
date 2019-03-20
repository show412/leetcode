# @param {Integer[]} heights
# @return {Integer}
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
