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

//删除
func (this *BinTree) delete(v int) {
	if this == nil {
		return
	}
	var parent *BinTree = nil
	for this != nil {
		if this.Data == v {
			break
		}
		parent = this
		if this.Data >= v {
			this = this.Left
		} else {
			this = this.Right
		}
	}
	//没有找到。
	if this == nil {
		return
	}
	var child *BinTree = nil
	//如果要删除的节点没有子节点
	if this.Right == nil && this.Left == nil {
		child = nil
		if parent.Left == this {
			parent.Left = child
		} else {
			parent.Right = child
		}
		return
	}
	//如果要删除的节点只有一个子节点。
	if (this.Right != nil && this.Left == nil) || (this.Left == nil || this.Right != nil) {
		if this.Left != nil {
			child = this.Left
		} else {
			child = this.Right
		}
		if parent.Right == this {
			parent.Right = child
		} else {
			parent.Left = child
		}
		return
	}
	//如果要删除的节点有两个子节点
}
