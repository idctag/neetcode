package arr

func ContainsDuplicate(nums []int) bool {
	mp := make(map[int]bool)
	for _, n := range nums {
		if _, ok := mp[n]; ok {
			return true
		}
		mp[n] = true
	}
	return false
}
