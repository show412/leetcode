/**
 * @param target: the target string
 * @param s:
 * @return: output all strings containing target  in s
 */
func getAns(target string, s []string) []string {
	// Write your code here
	if len(target) == 0 || len(s) == 0 {
		return []string{}
	}
	res := make([]string, 0)
	tLen := len(target)
	for i := 0; i < len(s); i++ {
		str := s[i]
		k := 0
		for j := 0; j < len(str); j++ {
			if string(str[j]) == string(target[k]) {
				if k == tLen-1 {
					res = append(res, str)
					break
				} else {
					k++
				}
			}

		}
	}
	return res
}
