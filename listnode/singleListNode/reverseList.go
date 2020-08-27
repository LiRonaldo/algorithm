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
	createListNode(3, head)
	sout(head)
	pre := reverse(head)
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

//普通反转,比较不好理解
func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	for cur != nil {
		cur.next, pre, cur = pre, cur, cur.next
	}
	return pre
}

//递归反转，还比较好理解。递归理解成栈，先进后出
func reverse(head *ListNode) *ListNode {
	if head.next == nil || head == nil {
		return head
	}
	//一直递归找到最后一个元素，则当前的head 就是倒数第二个元素。把第二个放到最后。然后，next=nil
	last := reverse(head.next)
	head.next.next = head
	head.next = nil
	return last
}