func textQueries(sentences []string, queries []string) [][]int {
  if sentences == nil || queries == nil {
    return nil
  } 
  res := make([][]int, len(sentences))
  if len(sentences) == 0 || len(queries) == 0 {
    return res
  } 
  // hast to store word to index
  m := make(map[string][]int)
  
  for k, v := range sentences {
    list := strings.Split(v, " ")
    for _, word := range list {
      if _, ok := m[word]; ok {
        _ = append(m[word], k)
      }else {
        m[word] = [k]
      }
    }
  }

  for _, v := range queries {
    queryArray = strings.Split(v, " ")
    for _, query := range queryArray {
      if _, ok := m[query]; ok {
        res = append(res, m[query])
      }
    }  
  }

  return res

}