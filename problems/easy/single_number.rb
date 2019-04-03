# @param {Integer[]} nums
# @return {Integer}
def single_number(nums)
  return 0 if(nums.length == 0)
  n = nums[0]
  i = 1
  while(i<nums.length)
    n = n ^ nums[i]
    i += 1
  end
  return n
end
