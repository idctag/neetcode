package stack

import "strconv"

func EvalRPN(tokens []string) int {
	stack := []int{}

	for _, token := range tokens {
		switch token {
		case "+":
			val := stack[len(stack)-2] + stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, val)
			continue
		case "-":
			val := stack[len(stack)-2] - stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, val)
			continue
		case "*":
			val := stack[len(stack)-2] * stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, val)
			continue
		case "/":
			val := stack[len(stack)-2] / stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, val)
			continue
		}

		num, _ := strconv.Atoi(token)
		stack = append(stack, num)
	}
	return stack[0]
}
