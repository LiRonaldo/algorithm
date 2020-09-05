package main

import (
	"algorithm/tree/queue"
	"fmt"
	"math"
)

/**
二叉树的前序，中序，后续，遍历。递归,层序，镜像，最大值
结点的度：结点拥有的子树的数目。
叶子：度为零的结点。
分支结点：度不为零的结点。
树的度：树中结点的最大的度。

层次：根结点的层次为1，其余结点的层次等于该结点的双亲结点的层次加1。
树的高度：树中结点的最大层次,从下往上数
数的深度:从上往下数.
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

/**
二叉树翻转，也就是镜像，用递归，从根节点开始先翻，在从左节点开始一直到递归出口，在翻转右一直到递归出口
*/
func reserve(t *Tree) *Tree {
	if t == nil || (t.right == nil && t.left == nil) {
		return nil
	}
	t.right, t.left = t.left, t.right
	reserve(t.left)
	reserve(t.right)
	return t
}

/**
求二叉树最大值 ,层序遍历，借助队列，找最大的.
*/
func findMax(t *Tree) int {
	max := t.data
	queue := &queue.Queue{}
	queue.Push(t)
	for !queue.IsEmpty() {
		t := queue.Pop()
		switch curT := t.(type) {
		case *Tree:
			if curT.data > max {
				max = curT.data
			}
			if curT.right != nil {
				queue.Push(curT.right)
			}
			if curT.left != nil {
				queue.Push(curT.left)
			}
		}
	}
	return max
}

/**
求树的最大深度，最大高度，最大层数，都是一样的，如果求某一个节点的话，要区分从上数，从下数。
*/
func maxDeep(t *Tree) float64 {
	if t == nil {
		return 0
	}
	return math.Max(1+maxDeep(t.left), 1+maxDeep(t.right))
}

/**
求树的最小深度，最小高度，最小层数，都是一样的，如果求某一个节点的话，要区分从上数，从下数。
*/
func minDeep(t *Tree) float64 {
	if t == nil {
		return 0
	}
	return math.Min(1+maxDeep(t.left), 1+maxDeep(t.right))
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
	t := reserve(root)
	rsult1 := SequenceTraversal(t)
	fmt.Println(rsult1)
	v := findMax(root)
	fmt.Println(v)
	deep := minDeep(t)
	fmt.Println(deep)
}
