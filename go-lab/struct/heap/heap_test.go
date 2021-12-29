package heap

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	arr := []int{10, 15, 25, 30, 56, 70}
	h := Heapify(arr)
	fmt.Println(h)
	fmt.Println(h.Insert(0))
	fmt.Println(h.Insert(100))
	fmt.Println(h.Insert(18))
	fmt.Println(h.Pop())
	fmt.Println(h)
}
