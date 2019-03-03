# @param {Integer[]} nums
# @return {Boolean}
def can_jump(nums)
  # greedy algorithm
  return false if(nums.nil? || nums.length == 0)
  farthest = nums[0]
  i = 1
  while(i<nums.length)
    farthest = i+nums[i] if(i <= farthest && i+nums[i]>= farthest)
    i += 1
  end
  farthest >= nums.length - 1
end
