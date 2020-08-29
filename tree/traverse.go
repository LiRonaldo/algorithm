package main

import "fmt"

/**
二叉树的前序，中序，后续，遍历。递归
*/
type Tree struct {
	data  interface{}
	left  *Tree
	right *Tree
}

/**
创建
*/
func creat(v int) *Tree {
	return &Tree{data: v}
}

/**
前序
*/
func pre(t *Tree) {
	if t == nil {
		return
	}
	fmt.Print(t.data)
	fmt.Print(" ")
	pre(t.left)
	pre(t.right)
}

/**
中序
*/
func mid(t *Tree) {
	if t == nil {
		return
	}
	mid(t.left)
	fmt.Print(t.data)
	fmt.Print(" ")
	mid(t.right)
}

/**
后续
*/
func after(t *Tree) {
	if t == nil {
		return
	}
	after(t.left)
	after(t.right)
	fmt.Print(t.data)
	fmt.Print(" ")
}
func main() {
	root := creat(4)
	root.left = creat(2)
	root.left.left = creat(1)
	root.left.right = creat(10)
	root.right = creat(6)
	root.right.left = creat(0)
	pre(root)
	fmt.Println()
	mid(root)
	fmt.Println()
	after(root)
}
