package stack

func DailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := []int{}
	for i, t := range temperatures {
		for len(stack) > 0 && t > temperatures[stack[len(stack)-1]] {
			stackInd := stack[len(stack)-1]
			res[stackInd] = i - stackInd
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return res
}
