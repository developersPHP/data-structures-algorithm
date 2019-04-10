package stack

import "fmt"

type node struct {
	next *node
	v    interface{}
}

type StackByList struct {
	topNode *node
}

func NewStackByList() *StackByList {
	return &StackByList{nil}
}

func (s *StackByList) IsEmpty() bool {
	if s.topNode.next == nil {
		return true
	}
	return false
}

func (s *StackByList) CreateStack() *StackByList {
	return s
}

func (s *StackByList) DisposeStack() {
	s.topNode = nil
}

func (s *StackByList) MakeEmpty() {
	if s == nil {
		return
	}
	for !s.IsEmpty() {
		s.Pop()
	}
}

func (s *StackByList) Push(e interface{}) {
	s.topNode = &node{next: s.topNode, v: e}
}

func (s *StackByList) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	v := s.topNode.v
	s.topNode = s.topNode.next
	return v
}

func (s *StackByList) Top() interface{} {
	return s.topNode.v
}

func Print(s *StackByList) {
	for !s.IsEmpty() {
		fmt.Println("stack element is %+v", s.Pop())
	}
}
