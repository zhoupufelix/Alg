package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}


func reverseKGroup(head *ListNode,k int)*ListNode{
	if head == nil {
		return head
	}
	a,b := head,head
	for i:=0;i<k;i++{
		//不足k个不需要反转
		if b == nil {
			return head
		}
		b = b.Next
	}
	//反转前k个
	newHead := reverse(a,b)
	//子问题
	a.Next = reverseKGroup(b,k)
	return newHead
}



//每隔k个反转链表
func reverse(head,end *ListNode)*ListNode{
	var pre *ListNode
	cur,next := head,head

	for cur!= end{
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func printListNode(head *ListNode){
	cur := head
	for cur != nil{
		fmt.Println(cur)
		cur = cur.Next
	}
}


func main(){
	head := new (ListNode)
	head.Val = 1

	listNode1 :=  new (ListNode)
	listNode1.Val = 2
	head.Next = listNode1

	listNode2 :=  new (ListNode)
	listNode2.Val = 3
	listNode1.Next = listNode2

	listNode3 :=  new (ListNode)
	listNode3.Val = 4
	listNode2.Next = listNode3

	listNode4 :=  new (ListNode)
	listNode4.Val = 5
	listNode3.Next = listNode4

	listNode5 :=  new (ListNode)
	listNode5.Val = 6
	listNode4.Next = listNode5

	newHead := reverseKGroup(head,2)
	printListNode(newHead)
}


