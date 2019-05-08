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
  j=1
  while (i<length) do
    break if (i+j>length)
    idx = s.index(r_s.slice(i,j))
    if !idx.nil?
      if check_palindro(s.slice(idx,j))
        longest_palindro = s.slice(idx,j)
        j += 1
      else
        i += 1
        j += 1
      end
    else
      i += 1
    end
  end
  return longest_palindro
end

def check_palindro(s)
  length1 = s.length
  i1 = 0
  j1 = length1-1
  flag = true
  while(i1<length1)
    break if i1 == j1 || j1 <0 || i1>j1
    if s[i1] != s[j1]
      flag = false
      break
    else
      i1 +=1
      j1 -=1
    end
  end
  return flag
end


# The best solution
def longest_palindrome(s)
    max_indexes = [0, 0]

    (0..(s.size - 1)).each do |idx|
        start_1, end_1 = expand_palindrome(idx, idx, s)
        # 这有隔一个的情况 比如 cbabc 中间只有一个a也是回文
        start_2, end_2 = expand_palindrome(idx, idx + 1, s)

        if end_1 - start_1 > max_indexes[1] - max_indexes[0]
            max_indexes = [start_1, end_1]
        end

        if end_2 - start_2 > max_indexes[1] - max_indexes[0]
            max_indexes = [start_2, end_2]
        end
    end

    s[max_indexes[0]..max_indexes[1]]
end

def expand_palindrome(start_idx, end_idx, s)
    while start_idx >= 0 && end_idx <= s.size - 1 && s[start_idx] == s[end_idx]
        start_idx -= 1
        end_idx += 1
    end

    return start_idx + 1, end_idx - 1
end
