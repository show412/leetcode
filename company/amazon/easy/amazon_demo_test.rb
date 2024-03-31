def cellCompete(states, days)
  # WRITE YOUR CODE HERE
  while days > 0
    days -= 1
    pre = states.clone
    (0..pre.length-1).each do |i|
      left = 0
      right = 0
      left = pre[i-1] if i != 0
      right = pre[i+1] if i != pre.length-1
      if left != right
        states[i] = 1
      else
        states[i] = 0
      end
    end
  end
  states
end
p cellCompete([1,0,0,0,0,1,0,0], 1)
