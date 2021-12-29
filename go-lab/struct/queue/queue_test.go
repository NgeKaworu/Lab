package queue

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	q := new(Queue)

	fmt.Println(q.Enqueue(1))
	fmt.Println(q.Dequeue())
	fmt.Println(q)
	// fmt.Println(q.Dequeue())s
	fmt.Println(q.Enqueue(2))
	fmt.Println(q.Enqueue(3))
}
