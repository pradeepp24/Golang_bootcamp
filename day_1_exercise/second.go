package main

import (
	"fmt"
)

type Tree struct {
	left  *Tree
	right *Tree
	val   byte
}

func (node *Tree) preorder() {
	if node == nil {
		return
	}
	fmt.Print(string(node.val))
	node.left.preorder()
	node.right.preorder()
}
func (node *Tree) postorder() {
	if node == nil {
		return
	}

	node.left.postorder()
	node.right.postorder()
	fmt.Print(string(node.val))
}

func main() {
	s := "a+b-c"
	root := &Tree{val: s[1]}
	root.left = &Tree{val: s[0]}
	root.right = &Tree{val: s[3]}
	root.right.left = &Tree{val: s[2]}
	root.right.right = &Tree{val: s[4]}

	root.preorder()
	fmt.Println()
	root.postorder()
}
