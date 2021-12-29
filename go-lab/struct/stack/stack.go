package stack

import (
	"log"
	"strconv"
)

// Stack 数据结构 栈
type Stack struct {
	stack []int
}

// Len 长度
func (s *Stack) Len() int {
	return len(s.stack)
}

func (s *Stack) String() (res string) {
	for _, v := range s.stack {
		res += strconv.Itoa(v)
	}
	return
}

// Push 压栈
func (s *Stack) Push(val int) *Stack {
	s.stack = append(s.stack, val)
	return s
}

// Pop 出栈
func (s *Stack) Pop() (res int) {
	if s.Len() < 1 {
		log.Fatal("Empty stack")
	}
	res = s.stack[s.Len()-1]
	s.stack = s.stack[:s.Len()-1]
	return
}
