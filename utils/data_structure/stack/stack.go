package stack

import "fmt"

type Stack struct {
	base []interface{}
}

func (s *Stack) Push(in interface{}) {
	s.base = append(s.base, in)
}

func (s *Stack) Top() interface{} {
	if len(s.base) <= 0 {
		panic("stack get fail: empty")
		return nil
	}

	return s.base[len(s.base)-1]
}

func (s *Stack) Pop() bool {
	if len(s.base) <= 0 {
		panic("stack get fail: empty")
		return false
	}

	s.base = s.base[:len(s.base)-1]

	return true
}

func (s *Stack) GetAndPop() (inf interface{}) {
	if len(s.base) <= 0 {
		panic("stack get fail: empty")
		return nil
	}

	inf = s.base[len(s.base)-1]
	s.base = s.base[:len(s.base)-1]

	return
}

func (s *Stack) Len() int {
	return len(s.base)
}

func (s *Stack) GetByIndex(index int) interface{} {
	if index > len(s.base)-1 {
		panic("stack get fail: empty")
		return nil
	}

	return s.base[index]
}

func (s *Stack) String() (str string) {
	for i := range s.base {
		str += fmt.Sprintf("%+v", s.base[i])
	}
	return
}
