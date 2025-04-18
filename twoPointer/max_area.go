package twopointer

func MaxArea(height []int) int {
	lidx, ridx, res := 0, len(height)-1, 0
	for lidx < ridx {
		area := min(height[lidx], height[ridx]) * (ridx - lidx)
		res = max(res, area)
		if height[lidx] > height[ridx] {
			ridx--
		} else if height[lidx] < height[ridx] {
			lidx++
		} else {
			ridx--
			lidx++
		}
	}
	return res
}
