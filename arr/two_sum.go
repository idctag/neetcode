package arr

func TwoSum(nums []int, target int) []int {
	count := make(map[int]int)
	for i, n := range nums {
		if v, ok := count[target-n]; ok {
			return []int{v, i}
		}
		count[n] = i
	}
	return []int{}
}
