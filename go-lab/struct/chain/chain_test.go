package chain

import (
	"fmt"
	"testing"
)

func TestChain(t *testing.T) {
	tests := []int{2, 8, 9, 7, 6, 5, 1, 0, -1}

	c := FromArr(tests)
	fmt.Println(c)
	fmt.Println(c.ToArr())

	fmt.Println(c.ToArr())
	fmt.Println(c.InsertH(-1100))
	fmt.Println(c.InsertT(110))
	fmt.Println(c.Insert(-1100, 111))
	fmt.Println(c.Insert(8, 112))
	fmt.Println(c.Insert(110, 113))
	fmt.Println(c.Update(110, 114))
	fmt.Println(c.Remove(-1100))
	// fmt.Println(c.Remove(110))
	fmt.Println(c.Remove(111))
	fmt.Println(c.Remove(112))
	fmt.Println(c.Remove(113))
	fmt.Println(c.Remove(114))
	fmt.Println(c.Find(5))

	fmt.Println(c.Swap(8, 9))
	fmt.Println(c.Swap(9, 7))
	fmt.Println(c.HeadSwap(-1))
	fmt.Println(c.Sort())
}
