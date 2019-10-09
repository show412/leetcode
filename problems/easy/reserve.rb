# https://leetcode.com/problems/reverse-integer/
# Given a 32-bit signed integer, reverse digits of an integer.

# Example 1:

# Input: 123
# Output: 321
# Example 2:

# Input: -123
# Output: -321
# Example 3:

# Input: 120
# Output: 21
# Note:
# Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.

def reverse(x)
  begin
    if x < 0
      ("-" + x.to_s[1..x.length].reverse).to_i
    else
      x.to_s.reverse.to_i
    end
  rescue Exception => e
     0
  end
end

# better for ruby
def reverse(x)
    a = x.to_s.reverse
    a[-1] == '-' ? -a.to_i : a.to_i
end

# best for C++ the though is same for any language
# Complexity Analysis

# Time Complexity: O(\log(x))O(log(x)). There are roughly \log_{10}(x)log
# ​10
# ​​ (x) digits in xx.
# Space Complexity: O(1)O(1).

class Solution {
public:
    int reverse(int x) {
        int rev = 0;
        while (x != 0) {
            int pop = x % 10;
            x /= 10;
            if (rev > INT_MAX/10 || (rev == INT_MAX / 10 && pop > 7)) return 0;
            if (rev < INT_MIN/10 || (rev == INT_MIN / 10 && pop < -8)) return 0;
            rev = rev * 10 + pop;
        }
        return rev;
    }
};
