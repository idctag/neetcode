package stack

import "slices"

func ValidParanthesis(s string) bool {
	stack := []rune{}
	for _, c := range s {
		switch c {
		case '{':
			stack = append(stack, c)
		case '(':
			stack = append(stack, c)
		case '[':
			stack = append(stack, c)

		case '}':
			if len(stack) < 1 {
				return false
			}
			top := stack[len(stack)-1]
			if top != '{' {
				return false
			} else {
				stack = slices.Delete(stack, len(stack)-1, len(stack))
			}
		case ')':
			if len(stack) < 1 {
				return false
			}
			top := stack[len(stack)-1]
			if top != '(' {
				return false
			} else {
				stack = slices.Delete(stack, len(stack)-1, len(stack))
			}
		case ']':
			if len(stack) < 1 {
				return false
			}
			top := stack[len(stack)-1]
			if top != '[' {
				return false
			} else {
				stack = slices.Delete(stack, len(stack)-1, len(stack))
			}
		}
	}
	return len(stack) == 0
}
