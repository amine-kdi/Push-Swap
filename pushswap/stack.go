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
