package stack

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{stack: []int{}, minStack: []int{}}
}

func (s *MinStack) Push(val int) {
	if len(s.stack) == 0 || val <= s.minStack[len(s.minStack)-1] {
		s.minStack = append(s.minStack, val)
	}
	s.stack = append(s.stack, val)
}

func (s *MinStack) Pop() {
	if len(s.stack) == 0 {
		return
	}
	if s.stack[len(s.stack)-1] == s.minStack[len(s.minStack)-1] {
		s.minStack = s.minStack[:len(s.minStack)-1]
	}
	s.stack = s.stack[:len(s.stack)-1]
}

func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1]
}

func (s *MinStack) GetMin() int {
	return s.minStack[len(s.minStack)-1]
}
