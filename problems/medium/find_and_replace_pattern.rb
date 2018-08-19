# You have a list of words and a pattern, and you want to know which words in words matches the pattern.

# A word matches the pattern if there exists a permutation of letters p so that after replacing every letter x in the pattern with p(x), we get the desired word.

# (Recall that a permutation of letters is a bijection from letters to letters: every letter maps to another letter, and no two letters map to the same letter.)

# Return a list of the words in words that match the given pattern.

# You may return the answer in any order.



# Example 1:

# Input: words = ["abc","deq","mee","aqq","dkd","ccc"], pattern = "abb"
# Output: ["mee","aqq"]
# Explanation: "mee" matches the pattern because there is a permutation {a -> m, b -> e, ...}.
# "ccc" does not match the pattern because {a -> c, b -> c, ...} is not a permutation,
# since a and b map to the same letter.


# Note:

# 1 <= words.length <= 50
# 1 <= pattern.length = words[i].length <= 20
# https://leetcode.com/problems/find-and-replace-pattern/description/
# Generate a hash which key is letter of words and values is the letter of words.
# When the according word is replaced by the pattern of 'pattern' is the same with pattern, it means the word is that we want to get

# @param {String[]} words
# @param {String} pattern
# @return {String[]}
def find_and_replace_pattern(words, pattern)
  length = words.length
  p_length = pattern.length
  pattern_array = pattern.split('').uniq
  pul=pattern_array.length
  out = []
  (0..length-1).each do |i|
    h = {}
    w=words[i].split('').uniq
    wul=w.length
    origin_word=words[i].clone
    if wul == pul
      (0..wul-1).each do |j|
        h[w[j]] = pattern_array[j]
      end
      (0..p_length-1).each do |z|
        if h[words[i][z]]
          words[i][z] = h[words[i][z]]
        end
      end
      out.push(origin_word) if words[i] == pattern
    end
  end
  out
end
