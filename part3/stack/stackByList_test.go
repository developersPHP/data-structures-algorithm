package stack

import "testing"

func TestStackByList_Push(t *testing.T) {
	s := NewStackByList()
	for i := 0; i < 3; i++ {
		s.Push(i)
	}
	Print(s)
}
