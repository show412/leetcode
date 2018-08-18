# You're given strings J representing the types of stones that are jewels, and S representing the stones you have.  Each character in S is a type of stone you have.  You want to know how many of the stones you have are also jewels.

# The letters in J are guaranteed distinct, and all characters in J and S are letters. Letters are case sensitive, so "a" is considered a different type of stone from "A".

# Example 1:

# Input: J = "aA", S = "aAAbbbb"
# Output: 3
# Example 2:

# Input: J = "z", S = "ZZ"
# Output: 0
# Note:

# S and J will consist of letters and have length at most 50.
# The characters in J are distinct.

# https://leetcode.com/problems/jewels-and-stones/description/

# @param {String} j
# @param {String} s
# @return {Integer}
def num_jewels_in_stones(j, s)
    return 0 if j == [] || s == [] || j.nil? || s.nil?
    i = 0
    count = 0
    length = s.length
    while(i<length) do
      count += 1 if j.include?(s[i])
      i+=1
    end
    count
end
