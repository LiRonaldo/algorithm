package main

import "fmt"

/**
双向链表
*/
type Element interface {
}
type Node struct {
	Data interface{}
	Next *Node
	Pre  *Node
}
type List struct {
	Head   *Node
	Length int
}

func Create() *List {
	node := new(Node)
	return &List{Head: node, Length: 0}
}
func (l *List) Len() (n int) {
	return l.Length
}
func (l *List) IsEmpty() bool {
	if l.Head.Next == nil && l.Head.Pre == nil {
		return true
	} else {
		return false
	}
}

//在链表尾部添加
func (l *List) Append(e Element) {
	node := &Node{Data: e, Next: nil, Pre: nil}
	p := l.Head
	if l.IsEmpty() {
		p.Next = node
		node.Pre = p
		return
	}
	for p.Next != nil {
		p = p.Next
	}
	p.Next = node
	node.Pre = p
	l.Length++
	return
}

//在链表头部添加
func (l *List) PreAppend(e Element) {
	node := &Node{Data: e, Next: nil, Pre: nil}
	p := l.Head
	if l.IsEmpty() {
		p.Next = node
		node.Pre = p
		return
	}
	nextp := p.Next.Next
	p.Next = node
	node.Pre = p
	nextp.Pre = node
	node.Next = nextp
	l.Length++
	return
}
func (l *List) Write() {
	for cur := l.Head; cur != nil; cur = cur.Next {
		fmt.Print(cur.Data)
	}
	fmt.Println("")
}
func (l *List) reverseList() {
	if l.IsEmpty() {
		return
	}
	p := l.Head.Next
	var pre *Node = nil
	var next *Node = nil
	for p != nil {
		next = p.Next
		p.Next = pre
		if next == nil {
			l.Head.Next = p
			p.Pre = l.Head
			break
		}
		pre = p
		p.Pre = next
		p = next
	}
}
func main() {
	l := Create()
	l.Append(3)
	l.Append(34)
	l.PreAppend(1)
	l.Write()
	l.reverseList()
	l.Write()
}
