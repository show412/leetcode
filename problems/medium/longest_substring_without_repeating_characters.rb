# Given a string, find the length of the longest substring without repeating characters.

# Examples:

# Given "abcabcbb", the answer is "abc", which the length is 3.

# Given "bbbbb", the answer is "b", with the length of 1.

# Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

# @param {String} s
# @return {Integer}
# https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
def length_of_longest_substring(s)
    length = s.length
    i = 0
    sub_str = []
    longest = 0
    while i<length
      index = sub_str.index(s[i])
      if !index.nil?
        longest = longest > sub_str.length ? longest : sub_str.length
        sub_str.slice!(0..index)
        sub_str.push(s[i])
      else
        sub_str.push(s[i])
      end
      i += 1
    end
    longest > sub_str.length ? longest : sub_str.length
end
