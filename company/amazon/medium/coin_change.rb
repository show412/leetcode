# // https://leetcode.com/problems/coin-change/
# /*
# You are given coins of different denominations
# and a total amount of money amount.
# Write a function to compute the fewest number of coins
# that you need to make up that amount.
# If that amount of money cannot be made up by any combination of the coins,
# return -1.

# Example 1:

# Input: coins = [1, 2, 5], amount = 11
# Output: 3
# Explanation: 11 = 5 + 5 + 1

# Example 2:

# Input: coins = [2], amount = 3
# Output: -1
# Note:
# You may assume that you have an infinite number of each kind of coin.
# */

# @param {Integer[]} coins
# @param {Integer} amount
# @return {Integer}
def coin_change(coins, amount)
  f = Array.new(amount+1)
  (0...amount+1).each {|i| f[i] = amount+1}
  f[0] = 0
  (0...amount+1).each do |i|
    (0...coins.length).each do |j|
      f[i] = [f[i], f[i-coins[j]]+1].min if i >= coins[j]
    end
  end
  return - 1 if f[amount] > amount
  f[amount]
end
