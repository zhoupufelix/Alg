package main

import "fmt"

//翻转一棵二叉树。
//
//示例：
//
//输入：
//
//	  4
//	/   \
// 2     7
/// \   / \
//1   3 6   9
//输出：
//
//    4
//  /   \
// 7     2
/// \   / \
//9   6 3   1


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode)*TreeNode{
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left

	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

func main(){

	leftLSub   := &TreeNode{Val: 1 }
	leftRSub   := &TreeNode{Val: 3 }
	leftNode   := &TreeNode{Val: 2,Left: leftLSub,Right: leftRSub}
	rightLSub  := &TreeNode{Val: 6 }
	rightRSub  := &TreeNode{Val: 9 }
	rightNode  := &TreeNode{Val: 7 ,Left:rightLSub,Right: rightRSub}
	root       := &TreeNode{Val: 3,Left:leftNode,Right: rightNode}

	invertTree(root)

	fmt.Println(root.Left)
}