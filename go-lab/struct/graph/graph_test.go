package graph

import (
	"fmt"
	"testing"
)

func TestGraph(t *testing.T) {
	v5 := NewVertex(5)
	v4 := NewVertex(4)
	v5.Connect(v4, 2, 200)
	v10 := NewVertex(10)
	v4.Connect(v10, 2, 700)
	v8 := NewVertex(8)
	v4.Connect(v8, 2, 300)
	v8.Connect(v5, 2, 30)
	v3 := NewVertex(3)
	v3.Connect(v5, 2, 20)
	v7 := NewVertex(7)
	v7.Connect(v8, 2, 3)
	v3.Connect(v7, 2, 40)
	v1 := NewVertex(1)
	v1.Connect(v7, 1, 6)
	v1.Connect(v5, 2, 4)

	g := &Graph{start: v5}
	fmt.Print(g)
	for _, v := range v5.FindAllPath(7) {
		fmt.Println(v)
	}
}
