package queue

import (
	"container/list"
	"fmt"
)

type QueueRecondList struct {
	l *list.List
}

func NewQueueRecondList() *QueueRecondList {
	l := list.New()
	return &QueueRecondList{l: l}
}

func (q *QueueRecondList) IsEmpty() bool {
	if q.l.Len() == 0 {
		return true
	}
	return false
}

//由于用go的链表实现队列，没有指定双向链表的长度，所以无法知道队列是否已经满了
func (q *QueueRecondList) IsFull() bool {
	return false
}

func (q *QueueRecondList) CreateQueue() *QueueRecondList {
	return q
}

func (q *QueueRecondList) DisposeQueue() {
	q.l.Init()
}

func (q *QueueRecondList) MakeEmpty() {
	for q.l.Len() != 0 {
		q.l.Front()
	}
}

func (q *QueueRecondList) Enqueue(e interface{}) {
	ele := list.Element{Value: e}
	q.l.PushBack(ele)
}

func (q *QueueRecondList) Dequeue() interface{} {
	e := q.l.Front()
	q.l.Remove(e)
	return e.Value
}

func Printv2(q *QueueRecondList) {
	for q.l.Len() != 0 {
		fmt.Println("the queue element is %v", q.Dequeue())
	}
}
