package main

import (
	"fmt"
	"math/rand"
)

/**
二叉查找数  二叉查找树中，每个节点的值都大于左子树节点的值，小于右子树节点的值
*/
/**
定义二叉查找树
*/
type BinTree struct {
	Data  int
	Left  *BinTree
	Right *BinTree
}

func main() {

	root := &BinTree{Data: 5}
	for i := 0; i < 5; i++ {
		root.insert(rand.Intn(100))
	}
	fmt.Println(root)
}
func (this *BinTree) insert(v int) {
	for this != nil {
		rootData := this.Data
		if v <= rootData {
			if this.Left == nil {
				this.Left = &BinTree{Data: v}
				break
			}
			this = this.Left
		} else {
			if this.Right == nil {
				this.Right = &BinTree{Data: v}
				break
			}
			this = this.Right
		}
	}
}
func (this *BinTree) find(v int) {
	if this == nil {
		return
	}
	for this != nil {
		if this.Data == v {
			return
		} else if this.Data > v {
			this = this.Left
		} else {
			this = this.Right
		}
	}
	return
}
