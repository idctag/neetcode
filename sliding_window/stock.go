package slidingwindow

func MaxProfit(prices []int) int {
	maxProfit, lidx := 0, 0

	for ridx := 1; ridx < len(prices); ridx++ {
		lprice := prices[lidx]
		rprice := prices[ridx]
		if lprice > rprice {
			lidx = ridx
			continue
		} else {
			maxProfit = max(maxProfit, rprice-lprice)
		}
	}
	return maxProfit
}
