package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}



func Sum(head1,head2 *ListNode)*ListNode{
	//哑节点
	dummy := &ListNode{}
	carry := 0 //进位
	//当前节点
	cur := dummy

	p := reverse(head1)
	q := reverse(head2)

	for p!= nil || q != nil{
		x := 0
		if p != nil {
			x = p.Val
		}
		y := 0
		if q != nil {
			y = q.Val
		}

		val := x+y+carry
		carry = val / 10
		val = val % 10

		cur.Next = &ListNode{Val: val}
		cur = cur.Next
		if (p != nil){
			p = p.Next
		}
		if (q != nil){
			q = q.Next
		}
	}

	if carry > 0 {
		cur.Next = &ListNode{Val: carry}
	}

	return reverse(dummy.Next)
}


func reverse(head *ListNode)*ListNode{
	var pre *ListNode
	cur := head
	for cur != nil{
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}





func main(){
	head := new (ListNode)
	head.Val = 9

	listNode1 :=  new (ListNode)
	listNode1.Val = 9
	head.Next = listNode1

	listNode2 :=  new (ListNode)
	listNode2.Val = 3
	listNode1.Next = listNode2

	listNode3 :=  new (ListNode)
	listNode3.Val = 4
	listNode2.Next = listNode3

	root := new(ListNode)
	root.Val = 6

	root1 := new(ListNode)
	root1.Val = 6
	root.Next = root1

	res := Sum(head,root)
	fmt.Println(res)
}