# // https://leetcode.com/problems/group-anagrams/
# /*
# Given an array of strings, group anagrams together.

# Example:

# Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
# Output:
# [
#   ["ate","eat","tea"],
#   ["nat","tan"],
#   ["bat"]
# ]
# Note:

# All inputs will be in lowercase.
# The order of your output does not matter.
# */
# @param {String[]} strs
# @return {String[][]}
def group_anagrams(strs)
  return [] if strs.empty?
  m = {}
  (0...strs.length).each do |i|
    # 初始化一个26长的数组 默认值是0
    count = Array.new(26, 0)
    str = strs[i]
    str.each_byte {|c| count[c-'a'.ord] += 1}
    sb = ""
    (0..25).each {|j| sb += count[j].to_s}
    m[sb] = [] if m.key?(sb) == false
    m[sb].push(str)
  end
  p m
  m.values
end
