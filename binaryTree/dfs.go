package main

import "fmt"

//给定一个二叉树，找出其最大深度。二叉树的深度为根节点到最远叶子节点的最长路径上的节点数
//示例
//给定二叉树 [3,9,20,null,null,15,7]

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}



func maxDepth(root *TreeNode)int{
	if root == nil{
		return 0
	}
	return Max(maxDepth(root.Left),maxDepth(root.Right))+1
} 



func Max(a,b int)int{
	if a > b{
		return a
	}
	return b
}

func main()  {
	leftNode   := &TreeNode{Val: 9 }
	rightLSub  := &TreeNode{Val: 15 }
	rightRSub  := &TreeNode{Val: 7 }
	rightNode  := &TreeNode{Val: 20 ,Left:rightLSub,Right: rightRSub}
	root       := &TreeNode{Val: 3,Left:leftNode,Right: rightNode}
	fmt.Println(maxDepth(root))
}
