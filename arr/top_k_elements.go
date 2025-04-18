package arr

func TopKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	bucket := make([][]int, len(nums)+1)

	for _, n := range nums {
		count[n]++
	}

	for k, v := range count {
		bucket[v] = append(bucket[v], k)
	}
	var res []int

	for i := len(bucket) - 1; i >= 0; i-- {
		for _, v := range bucket[i] {
			res = append(res, v)
			if len(res) == k {
				return res
			}
		}
	}

	return []int{}
}
