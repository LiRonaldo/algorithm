package main

import (
	"algorithm/tree/queue"
	"fmt"
)

/**
二叉树的前序，中序，后续，遍历。递归
*/
type Tree struct {
	data  int
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

/**
层序遍历,也就是二叉树的广度优先遍历（BFS）
借用模拟队列,//和java不一样的地方是interface相当于java 的object 但是，转化的时候还不能直接转，也不能强转，
必须判断类型
*/
func SequenceTraversal(t *Tree) []int {
	result := []int{}
	queue := &queue.Queue{}
	if t != nil {
		queue.Push(t)
	}
	for !queue.IsEmpty() {
		t := queue.Pop()
		switch curT := t.(type) {
		case *Tree:
			result = append(result, curT.data)
			if curT.left != nil {
				queue.Push(curT.left)
			}
			if curT.right != nil {
				queue.Push(curT.right)
			}
		}
	}
	return result
}
func main() {
	root := creat(4)
	root.left = creat(2)
	root.left.left = creat(1)
	root.left.right = creat(10)
	root.right = creat(6)
	root.right.left = creat(0)
	//pre(root)
	//fmt.Println()
	//mid(root)
	//fmt.Println()
	//after(root)
	rsult := SequenceTraversal(root)
	fmt.Println(rsult)
}
