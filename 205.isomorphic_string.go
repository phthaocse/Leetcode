func isIsomorphic(s string, t string) bool {
	lenStr := len(s)
	mapS := map[byte]int{}
	mapT := map[byte]int{}
	for i := 0; i < lenStr; i++ {
		if _, ok := mapS[s[i]]; !ok {
			mapS[s[i]] = i
		}
		if _, ok := mapT[t[i]]; !ok {
			mapT[t[i]] = i
		}
	}
	for i := 0; i < lenStr; i++ {
		if mapS[s[i]] == mapT[t[i]] {
			continue
		}
		return false
	}
	return true
}
