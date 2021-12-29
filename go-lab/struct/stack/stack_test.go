package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	s := new(Stack)
	fmt.Println(s.Push(1))
	fmt.Println(s.Push(2))
	fmt.Println(s.Push(3))
	fmt.Println(s.Pop())
	fmt.Println(s)

}
