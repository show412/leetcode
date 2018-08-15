# Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

# Example 1:

# Input: "babad"
# Output: "bab"
# Note: "aba" is also a valid answer.
# Example 2:

# Input: "cbbd"
# Output: "bb"
# https://leetcode.com/problems/longest-palindromic-substring/description/
# @param {String} s
# @return {String}
def longest_palindrome(s)
  longest_palindro = s[0] || ""
  r_s = s.reverse
  length = r_s.length
  i=0
  j=2
  while (i<length) do
    break if (j+i > length)
    idx = s.index(r_s.slice(i,j))
    if !idx.nil?
      longest_palindro = s.slice(idx,j)
      j += 1
    else
      i += 1
    end
  end
  return longest_palindro
end
