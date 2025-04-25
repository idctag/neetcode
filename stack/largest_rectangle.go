package stack

func LargestRectangleArea(heights []int) int {
	stack := [][2]int{} // index, height
	maxArea := 0
	for i, h := range heights {
		newIdx := i
		for len(stack) > 0 && stack[len(stack)-1][1] > h {
			p := stack[len(stack)-1]
			area := (i - p[0]) * p[1]
			maxArea = max(area, maxArea)
			newIdx = p[0]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, [2]int{newIdx, h})
	}
	for i := len(stack) - 1; i >= 0; i-- {
		p := stack[i]
		area := (len(heights) - p[0]) * p[1]
		maxArea = max(area, maxArea)
	}
	return maxArea
}
