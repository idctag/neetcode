package slidingwindow

func LongestSubstring(s string) int {
	mp := make(map[byte]int)
	l, res := 0, 0
	for r := range len(s) {
		if idx, found := mp[s[r]]; found {
			l = max(idx+1, l)
		}
		mp[s[r]] = r
		res = max(res, r-l+1)
	}
	return res
}
