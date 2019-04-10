package queue

import "testing"

func TestQueueRecondList_Dequeue(t *testing.T) {
	q := NewQueueRecondList()
	for i := 0; i < 3; i++ {
		q.Enqueue(i)
	}
	Printv2(q)
}
