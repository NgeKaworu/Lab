package queue

import (
	"log"
	"strconv"
)

// Queue 队列
type Queue struct {
	queue []int
}

// Len 长度
func (q *Queue) Len() int {
	return len(q.queue)
}

func (q *Queue) String() (s string) {
	for _, v := range q.queue {
		s += strconv.Itoa(v)
	}
	return
}

// Enqueue 进队
func (q *Queue) Enqueue(val int) *Queue {
	q.queue = append(q.queue, val)
	return q
}

// Dequeue 出队
func (q *Queue) Dequeue() (res int) {
	if q.Len() < 1 {
		log.Fatal("Emtpy queue")
	}
	res = q.queue[0]
	q.queue = q.queue[1:]
	return
}
