# @param {Integer[]} nums
# @return {Integer[]}
def majority_element(nums)
  #     quick sort first then travers the array
  # quick sort Time complexity is O(logn) space complexity is O(1)
  return [] if nums.nil? || nums.empty?
  nums.sort!
  length = nums.length
  size = length / 3
  result = []
  i = 0
  j = 0
  count = 1
  while(j < length)
    j += 1
    if nums[i] == nums[j]
      count += 1
    else
      result.push(nums[i]) if count > size
      i = j
      count = 1
    end
  end
  result
end