package stack

import "testing"

func TestStackByArray_Push(t *testing.T) {
	s := CreateStack(12)
	for i := 0; i < 12; i++ {
		s.Push(i)
	}
	Printv2(s)
}
