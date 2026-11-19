package pushswap

import "errors"

var ErrEmptyStack = errors.New("pile vide")

type Stack struct {
	items []int
}

func (s *Stack) Push(v int) {
	s.items = append(s.items, v)
}

func (s *Stack) Pop() (int, error) {
	if len(s.items) == 0 {
		return 0, ErrEmptyStack
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top, nil
}

func (s *Stack) Top() (int, error) {
	if len(s.items) == 0 {
		return 0, ErrEmptyStack
	}
	return s.items[len(s.items)-1], nil
}

func (s *Stack) Len() int {
	return len(s.items)
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Get(i int) int {
	return s.items[len(s.items)-1-i]
}

func (s *Stack) Copy() Stack {
	nouveau := make([]int, len(s.items))
	copy(nouveau, s.items)
	return Stack{items: nouveau}
}

func (s *Stack) Swap() {
	n := len(s.items)
	if n < 2 {
		return
	}
	temp := s.items[n-1]
	s.items[n-1] = s.items[n-2]
	s.items[n-2] = temp
}

func (s *Stack) Rotate() {
	n := len(s.items)
	if n < 2 {
		return
	}
	haut := s.items[n-1]
	for i := n - 1; i > 0; i-- {
		s.items[i] = s.items[i-1]
	}
	s.items[0] = haut
}

func (s *Stack) ReverseRotate() {
	n := len(s.items)
	if n < 2 {
		return
	}
	bas := s.items[0]
	for i := 0; i < n-1; i++ {
		s.items[i] = s.items[i+1]
	}
	s.items[n-1] = bas
}

func (s *Stack) IsSortedFromTop() bool {
	for i := 0; i < len(s.items)-1; i++ {
		if s.items[i] < s.items[i+1] {
			return false
		}
	}
	return true
}
