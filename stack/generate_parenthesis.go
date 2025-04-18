package stack

func GenerateParenthesis(n int) []string {
	res := []string{}
	stack := []byte{}
	addParenthesis(0, 0, stack, n, &res)
	return res
}

func addParenthesis(openN int, closeN int, stack []byte, n int, res *[]string) {
	if openN == closeN && openN == n {
		*res = append(*res, string(stack))
		return
	}
	if openN < n {
		newStack := stack
		newStack = append(newStack, '(')
		addParenthesis(openN+1, closeN, newStack, n, res)
	}
	if openN > closeN {
		newStack := stack
		newStack = append(newStack, ')')
		addParenthesis(openN, closeN+1, newStack, n, res)
	}
}
