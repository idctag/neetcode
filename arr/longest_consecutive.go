package arr

func LongestConsecutive(nums []int) int {
	longest := 0
	count := make(map[int]int)
	for _, n := range nums {
		if count[n] == 0 {
			left := count[n-1]
			right := count[n+1]
			sum := left + right + 1

			count[n-left] = sum
			count[n+right] = sum
			count[n] = sum
			longest = max(longest, sum)
		}
	}
	return longest
}
