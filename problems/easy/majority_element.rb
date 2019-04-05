# @param {Integer[]} nums
# @return {Integer}
def majority_element(nums)
  return nil if nums.length == 0 || nums.nil?
  i = 0
  candidate = 0
  count = 0
  while(i<nums.length)
    candidate = nums[i] if count == 0

    if(candidate == nums[i])
      count += 1
    else
      count -= 1
    end
    i += 1
  end
  return candidate
end
