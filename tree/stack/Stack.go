package stack

type Stack struct {
	items []interface{}
}

/**
判断是否为空
*/
func (s *Stack) IsEmpty() bool {
	if len(s.items) == 0 {
		return true
	} else {
		return false
	}
}

/**
压栈
*/
func (s *Stack) Push(v interface{}) {
	s.items = append(s.items, v)
}

/**
出栈
*/
func (s *Stack) Pop() interface{} {
	item := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]
	return item
}
