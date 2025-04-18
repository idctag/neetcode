package arr

func GroupAnagrams(strs []string) [][]string {
	res := make(map[[26]int][]string)
	for _, str := range strs {
		var count [26]int
		for _, char := range str {
			count[char-'a']++
		}
		res[count] = append(res[count], str)
	}
	var result [][]string
	for _, v := range res {
		result = append(result, v)
	}
	return result
}
