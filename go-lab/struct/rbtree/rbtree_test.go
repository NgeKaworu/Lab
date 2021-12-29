package rbtree

import (
	"fmt"
	"testing"
)

func TestRBTree(t *testing.T) {
	tree := new(RBTree)
	fmt.Println(tree.Insert(7))
	fmt.Println(tree.Insert(3))
	fmt.Println(tree.Insert(4))
	fmt.Println(tree.Insert(1))
	fmt.Println(tree.Insert(9))
	fmt.Println(tree.Insert(10))
	fmt.Println(tree.Insert(8))

	fmt.Println(tree.Remove(4))
	fmt.Println(tree)
	fmt.Println(tree.Remove(7))
	fmt.Println(tree)

}
