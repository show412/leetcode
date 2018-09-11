# The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

# P   A   H   N
# A P L S I I G
# Y   I   R
# And then read line by line: "PAHNAPLSIIGYIR"

# Write the code that will take a string and make this conversion given a number of rows:

# string convert(string s, int numRows);
# Example 1:

# Input: s = "PAYPALISHIRING", numRows = 3
# Output: "PAHNAPLSIIGYIR"
# Example 2:

# Input: s = "PAYPALISHIRING", numRows = 4
# Output: "PINALSIGYAHRPI"
# Explanation:

# P     I    N
# A   L S  I G
# Y A   H R
# P     I
# https://leetcode.com/problems/zigzag-conversion/description/


# @param {String} s
# @param {Integer} num_rows
# @return {String}
def convert(s, num_rows)
  length = s.length
  return s if  num_rows == 1
  i = 0
  round = 0
  c_array = []
  s_array = []
  cs=""
  f = num_rows - 2
  while( i< length) do

    # if (c_array.length + 1 + round%num_rows) %num_rows == 0 || c_array.length == 0
    if (c_array.length + 1) % (num_rows -1) == 1 || c_array.length == 0
      s_array << s[i]
    else
      j = 0
      while(j < num_rows) do
        if j ==  f
          s_array[j] = s[i]
        else
          s_array[j] = ""
        end
        j += 1
      end
    end
    if ((c_array.length + 1) % (num_rows -1) != 1) && c_array.length != 0
      f -= 1
      f = num_rows - 2 if f == 0
      f = num_rows - 1 if f < 0
    end

    if s_array.length == num_rows
      c_array << s_array
      s_array = []
    end

    # round += 1 if c_array.length > 0  && c_array.length%num_rows == 0
    i+=1
  end
  c_array << s_array if s_array.length > 0
  # p c_array
  j=0
  while(j<num_rows) do
    i = 0
    while(i<c_array.length) do
      cs += c_array[i][j] || ""
      i+=1
    end
    j +=1
  end
  cs
end
