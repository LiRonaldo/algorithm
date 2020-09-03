package queue

/**
模拟队列
*/

type Queue struct {
	items []interface{}
}

//判断队列是否为空
func (s *Queue) IsEmpty() bool {
	if len(s.items) == 0 {
		return true
	} else {
		return false
	}
}

//入队
func (s *Queue) Push(t interface{}) {
	s.items = append(s.items, t)
}

//出队
func (s *Queue) Pop() interface{} {
	item := s.items[0]
	s.items = s.items[1:len(s.items)]
	return item
}
