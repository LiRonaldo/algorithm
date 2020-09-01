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
	if this.Left != nil && this.Right != nil {
		min := this.Right
		minp := this
		//从右节点开始，一直找左边的，找到左边最后一个节点(为什么？思考二叉查找树的特点)
		for min.Left != nil {
			minp = min
			min = min.Left
		}
		//将左边找到最小的值赋给要删除的点。
		this.Data = min.Data
		//如果左边你最小的节点没有子节点的话，代码到上边就结束了。
		//下边的代码是将左边的最小节点之后的节点作出相应的调整。
		if minp.Right != min {
			//因为一直走的是左边，所以只会进到这个这个if，将左边最小的节点的右节点整个赋值给父节点的左节点（因为左节点赋给了要删除的节点了，所以空掉了）
			minp.Left = min.Right
		} else {
			minp.Right = min.Right
		}
	}
}
