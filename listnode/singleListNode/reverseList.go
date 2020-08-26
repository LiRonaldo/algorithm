package main

import "fmt"

/**
单链表反转
*/
type ListNode struct {
	data interface{}
	next *ListNode
}

func main() {
	head := &ListNode{}
	createListNode(10, head)
	sout(head)
	pre := reverseList(head)
	sout(pre)
}
func createListNode(n int, node *ListNode) {
	cur := node
	for i := 1; i <= n; i++ {
		cur.next = &ListNode{}
		cur.data = i
		cur = cur.next
	}
}

func sout(head *ListNode) {
	for cur := head; cur != nil; cur = cur.next {
		fmt.Print(cur.data)
	}
	fmt.Println("")
}

func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	for cur != nil {
		cur.next, pre, cur = pre, cur, cur.next
	}
	return pre
}
