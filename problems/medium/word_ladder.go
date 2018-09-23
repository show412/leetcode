func ladderLength(beginWord string, endWord string, wordList []string) int {
  for _, v := range wordList {
    if v ==  endWord {
      return 0
    }
	}
  if len(beginWord) == 1 {
    return 2
  }
	queue := []string{beginWord}
	visited := make(map[string]bool)
	distance := 1
	for len(queue) != 0 {
		current_queue := queue
		q_length := len(current_queue)
    queue := []string{}
		for  i := 0; i<q_length-1; i++  {
      new_current_queue := current_queue[:q_length-i]
			word := current_queue[q_length - i]
			length := strings.Count(word,"")-1
			for j := 0; j<length-1; j++ {
        char_map := []string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}
        for  _, z := range char_map{
          tmp1 := []byte(word)
          tmp1[j] = z
          tmp := string(tmp1)
					if tmp == endWord && wordList[tmp] {
						distance++
						return distance
					}
					if wordList[tmp] && visited[tmp] != true && tmp != word {
            queue := queue.append(tmp)
						visited[word] = true
					}
				}
			}
		}
    if len(queue) != 0 {
			distance++
		}
	}
  if distance == 1 {
    return 0
  }
  return distance
}