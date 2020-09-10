package main

import "fmt"

/**
单链表反转
*/
type ListNode struct {
	data int
	next *ListNode
}

func main() {
	head := &ListNode{data: 1}
	//head1 := &ListNode{}
	createListNode(10, head)
	//createListNode(2, head1)
	//sout(head)
	//pre := reverse(head)
	//sout(pre)
	//delete(head, 1)
	//node := merger(head, head1)
	//sout(node)
	sout(head)
	reciprocalDel(head, 10)
	sout(head)

}
func createListNode(n int, node *ListNode) {
	cur := node
	for i := 1; i <= n; i++ {
		cur.next = &ListNode{data: i + 1}
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
	//反转核心代码，
	head.next.next = head
	//防止产生环
	head.next = nil
	return last
}

/**
删除某个节点.
*/
func delete(node *ListNode, v interface{}) {
	var preNode *ListNode = nil
	for node.next != nil {
		preNode = node
		node = node.next
		if node.data == v {
			//是最后一个节点
			if node.next == nil {
				preNode.next = nil
			} else {
				//不是最后一个节点
				preNode.next = node.next
			}
		}
	}
}

/**
判断链表是否有环，两个指针，一个块，一个慢，如果块的追上慢的，或者，快的在慢点后边。
*/
func circulation(t *ListNode) bool {
	slow := t
	fast := t
	for {
		if slow == nil || fast == nil {
			return false
		} else if fast == slow || fast.next == slow { //fast在slow后边，环也就是slow。
			return true
		} else {
			slow = slow.next
			fast = fast.next.next
		}
	}
}

/**
合并两个有序，单链表。递归。
*/
func merger(t1 *ListNode, t2 *ListNode) *ListNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	var mergerNode *ListNode = nil
	if t1.data > t2.data {
		mergerNode = t2
		mergerNode.next = merger(t1, t2.next)
	} else {
		mergerNode = t1
		mergerNode.next = merger(t1.next, t2)
	}
	return mergerNode
}

/**
删除倒数第k个元素。利用两个指针，i指针先走k，j指针在开始走，当i走到头的时候，j后边的第一个元素就是要删除的。
*/
func reciprocalDel(t *ListNode, v int) {
	if t == nil {
		return
	}
	fast := t
	slow := t
	var i int = 1
	for i <= v {
		i++
		fast = fast.next
	}
	for fast.next != nil {
		fast = fast.next
		slow = slow.next
	}
	slow.next = slow.next.next
}
