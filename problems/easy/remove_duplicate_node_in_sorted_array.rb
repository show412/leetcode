# https://leetcode.com/problems/remove-duplicates-from-sorted-array/submissions/
# @param {Integer[]} nums
# @return {Integer}
def remove_duplicates(nums)
  return nil if nums.nil?
  length = nums.length
  return length if length == 0
  i = 1
  cur_index = 1
  cur_value = nums[0]
  while i<length
    if nums[i] == cur_value
      i += 1
      next
    else
      cur_value = nums[i]
      nums[cur_index] = nums[i]
      cur_index += 1
    end
    i += 1
  end
  # p nums
  return cur_index
end