package stack

func DailyTemperatures(temperatures []int) []int {
	res := []int{}
	stack := []map[int]int{}
	for i, t := range temperatures {
		for stack[len(stack)-1] {
			temp, indx := stack[len(stack)-1]
		}
	}
	return res
}
