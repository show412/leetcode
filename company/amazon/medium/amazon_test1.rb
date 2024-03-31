def topNCompetitors(numCompetitors, topNCompetitors, competitors, numReviews, reviews)
  # WRITE YOUR CODE HERE
    # h record the word in review occur number in reviews
    h = Hash.new(0)
    (0...numReviews).each do |i|
      review = reviews[i]
      # split by space and punctuation
      reviewWord = review.split
      reviewWord.each { |word| h[word] += 1}
    end
    # p h
    # find the competitors occur number to record into h1
    h1 = Hash.new(0)
    (0...numCompetitors).each { |j| h1[competitors[j]] = h[competitors[j]] } 
    # p h1
    # h1 sort by value and key
    sortCompetiter = h1.sort do |a, b|
      if a[1] < b[1]
        b[1] <=> a[1]
      elsif a[1] == b[1]
        a[0] <=> b[0]
      else
        0
      end
    end
    # generated the result
    res = []
    while topNCompetitors > 0
      res.push(sortCompetiter.shift[0])
      topNCompetitors -= 1
    end
    res
  end
  p topNCompetitors(5,2,["a","b","c","d","e"],3,["d ssasfsa", "d ad dsa d qwesa", "hfd kshfjd dsahfjfads dsa"])