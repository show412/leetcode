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
p coin_change([1,2,5], 11)
