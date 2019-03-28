# @param {Integer[]} nums
# @return {Integer}
# https://leetcode.com/problems/longest-consecutive-sequence/submissions/
# use hash data struct to store the nums if or not appear, and it only traverse the nums
# twice, so the time complexity is O(2n) is O(n)
def longest_consecutive(nums)
  hash_map = Hash[nums.collect{|n| [n,1]}]
  results = []
  length = nums.length
  i = 0
  while(i<length)
    if(!hash_map[nums[i]-1])
      seq = nums[i]
      result = 0
      while(hash_map[seq])
        seq += 1
        result += 1
      end
      results << result
    end
    i += 1
  end
  results.max || 0
end
