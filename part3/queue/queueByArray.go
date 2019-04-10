package queue

import "fmt"

type QueueRecord struct {
	capacity    int
	front       int
	rear        int
	size        int
	elementType []interface{}
}

func NewQueueRecord() *QueueRecord {
	return &QueueRecord{}
}

func (q *QueueRecord) IsEmpty() bool {
	if q.size == 0 {
		return true
	}
	return false
}

func (q *QueueRecord) IsFull() bool {
	if q.size == q.capacity {
		return true
	}
	return false
}

func (q *QueueRecord) CreateQueue(maxElements int) *QueueRecord {
	q.capacity = maxElements
	q.size = 0
	q.rear = 0
	q.front = 0
	q.elementType = make([]interface{}, maxElements)
	return q
}

func (q *QueueRecord) DisposeQueue() {
	q.size = 0
	q.front = 0
	q.elementType = nil
	q.capacity = 0
	q.rear = 0
}

func (q *QueueRecord) MakeEmpty() {
	for q.size != 0 {
		q.Dequeue()
	}
}

func (q *QueueRecord) Enqueue(element interface{}) bool {
	if q.rear == q.capacity {
		if q.front == 0 {
			return false //表示整个队列已经满了
		}
		//需要搬运数据
		for i := q.front; i < q.rear; i++ {
			q.elementType[i-q.front] = q.elementType[i]
		}
		q.rear -= q.front
		q.front = 0
	}

	q.elementType[q.rear] = element
	q.rear += 1
	q.size += 1
	return true

}

func (q *QueueRecord) Dequeue() interface{} {
	if q.front == q.rear {
		return nil
	}
	ret := q.elementType[q.front]
	q.front += 1
	q.size -= 1
	return ret
}

func PrintQueue(q *QueueRecord) {
	for q.size != 0 {
		fmt.Printf("队列里内容为%+v", q.Dequeue())
	}
}

//func (q *QueueRecord) FrontAndDequeue() int {
//
//}
