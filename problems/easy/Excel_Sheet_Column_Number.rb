# Given a column title as appear in an Excel sheet, return its corresponding column number.
# 
# For example:
# 
#     A -> 1
#     B -> 2
#     C -> 3
#     ...
#     Z -> 26
#     AA -> 27
#     AB -> 28 
#     ...
# Example 1:
# 
# Input: "A"
# Output: 1
# Example 2:
# 
# Input: "AB"
# Output: 28
# Example 3:
# 
# Input: "ZY"
# Output: 701
# https://leetcode.com/problems/excel-sheet-column-number/description/

# @param {String} s
# @return {Integer}
def title_to_number(s)
  return 0 if s.nil?
  length = s.length
  i = 0
  num = 0
  while(i<length)
    num += (s[i].ord - 64)* (26**(length - i - 1))
    i += 1
  end
  num
end
