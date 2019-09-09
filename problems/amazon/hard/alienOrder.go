// https://leetcode.com/problems/alien-dictionary/
// 拓扑排序实现
/* test case
["ri","xz","qxf","jhsguaw","dztqrbwbm","dhdqfb","jdv","fcgfsilnb","ooby"]

*/
func alienOrder(words []string) string {
	// record all char
	charArray := make(map[byte]bool, 0)
	// record in degree is not 0 at first
	// 注意这里的key要是一个数组 因为一个点的入度的边不止一个
	startM := make(map[byte][]byte, 0)
	// record degree
	in := make([]int, 256)
	q := make([]byte, 0)
	res := make([]byte, 0)
	for _, word := range words {
		for i := 0; i < len(word); i++ {
			charArray[word[i]] = true
		}
	}
	// generate the sequence for every char in every word
	// 这个用来记每个点的拓扑顺序， 因为前后顺序是根据word在数组中的顺序确定的 一旦找到一个就可以break
	// 因为这个时候这个顺序已经确定了, 如果真的有不能确定拓扑顺序的点 那就是一个孤立的点 认为入度为0
	// 可以放到开始
	for i := 0; i < len(words)-1; i++ {
		min := min(len(words[i]), len(words[i+1]))
		j := 0
		for j < min {
			if words[i][j] != words[i+1][j] {
				startM[words[i][j]] = append(startM[words[i][j]], words[i+1][j])
				break
			}
			j++
		}
		if j == min && len(words[i]) > len(words[i+1]) {
			return ""
		}
	}
	// generate in degree
	for _, arr := range startM {
		for _, v := range arr {
			in[v-'a']++
		}
	}
	// push those char in degree is 0 into q and res
	for v, _ := range charArray {
		if in[v-'a'] == 0 {
			q = append(q, v)
			res = append(res, v)
		}
	}
	// start to topology sort, to remove one node and edge and minus the degree
	// put every node which degree is 0 into q and res
	for len(q) != 0 {
		c := q[0]
		q = q[1:]
		if _, ok := startM[c]; ok {
			arr := startM[c]
			for _, v := range arr {
				in[v-'a']--
				if in[v-'a'] == 0 {
					q = append(q, v)
					res = append(res, v)
				}
			}
		}
	}
	// 如果出来的结果和记的char是一样个数 说明没有环
	// 有环的话res会比charArray的长度长
	// 或者也可以去遍历in数组 看看最后有没有还大于0的key
	if len(res) == len(charArray) {
		return string(res)
	}

	return ""
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
