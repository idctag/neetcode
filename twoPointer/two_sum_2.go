package twopointer

func TwoSum2(numbers []int, target int) []int {
	l := 0
	r := len(numbers) - 1
	for l < r {
		curr := numbers[l] + numbers[r]
		if target == curr {
			return []int{l + 1, r + 1}
		}
		if target < curr {
			r--
			continue
		}
		if target > curr {
			l++
			continue
		}
	}
	return []int{0, 0}
}
