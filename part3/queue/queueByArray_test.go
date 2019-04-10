package queue

import (
	"fmt"
	"testing"
)

func TestQueueRecord_Dequeue(t *testing.T) {
	//创建一个长度为12的队列
	q := NewQueueRecord().CreateQueue(12)
	fmt.Printf("队列对象为%+v", q)
	for i := 0; i < 3; i++ {
		e := i + 12
		q.Enqueue(e)
	}
	PrintQueue(q)
	//fmt.Printf("队列对象为%+v", q)
	//fmt.Printf("出队列%+v", q.Dequeue())
	//fmt.Printf("出队列%+v", q.Dequeue())
	//fmt.Printf("出队列%+v", q.Dequeue())
	//
	//fmt.Printf("队列对象为%+v", q)

}
