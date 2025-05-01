package slidingwindow

func MaxProfit(prices []int) int {
	l, res := 0, 0
	for r := 1; r < len(prices); r++ {
		if prices[r] < prices[l] {
			l = r
			continue
		} else {
			res = max((prices[r] - prices[l]), res)
		}
	}
	return res
}
