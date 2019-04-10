package stack

import "fmt"

type StackByArray struct {
	top  int
	data []interface{}
}

func CreateStack(capacity int) *StackByArray {
	return &StackByArray{top: -1, data: make([]interface{}, capacity)}
}

func (this *StackByArray) IsEmpty() bool {
	if this.top < 0 {
		return true
	}

	return false

}

func (this *StackByArray) IsFull() bool {
	if this.top == len(this.data) {
		return true
	}
	return false
}

func (this *StackByArray) DisposeStack() {
	this.top = -1
	this.data = nil
}

func (this *StackByArray) MakeEmpty() {
	for !this.IsEmpty() {
		this.Pop()
	}
}

func (this *StackByArray) Push(e interface{}) {
	if this.IsEmpty() {
		this.top = 0
	} else {
		this.top += 1
	}
	if this.top > len(this.data)-1 {
		this.data = append(this.data, e)
	} else {
		this.data[this.top] = e
	}
}

func (this *StackByArray) Pop() interface{} {
	if this.top < 0 {
		return nil
	}
	v := this.data[this.top]
	this.top -= 1
	return v
}

func (this *StackByArray) Top() interface{} {
	if this.top < 0 {
		return nil
	}
	return this.data[this.top]
}

func Printv2(s *StackByArray) {
	for !s.IsEmpty() {
		fmt.Printf("stack element is %v", s.Pop())
	}
}
