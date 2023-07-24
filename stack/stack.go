package stack

type Stack struct {
	contents []int
	pointer  int
}

func (s *Stack) Push(i int) {
	s.contents = append(s.contents, i)
	s.pointer++
}

func (s Stack) Peek() int {
	if len(s.contents) == 0 {
		return -1
	}

	return s.contents[s.pointer-1]
}

func (s *Stack) Pop() {
	s.pointer--
}
