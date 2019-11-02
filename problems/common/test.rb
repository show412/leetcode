def most_common_word(paragraph, banned)
  m = {}
  # store banned into a hash for performance
  banned.each {|b| m[b] = true}
  countMap = {}
  word = ""
  (0...paragraph.length).each do |i|
    c = paragraph[i]
    if c =~ /[A-Z]/ || c =~ /[a-z]/
      word += c
    elsif word == ""
      next
    else
      if m[word.downcase] == true
        word = ""
        next
      end
      if countMap.has_key?(word.downcase)
        countMap[word.downcase] += 1
      else
        countMap[word.downcase] = 1
      end
      word = ""
    end
  end
  if word != ""
    if countMap.has_key?(word.downcase)
      countMap[word.downcase] += 1
    else
      countMap[word.downcase] = 1
    end
  end
  countMap.max_by{|k, v| v}[0]
end
p most_common_word("Bob", [])
