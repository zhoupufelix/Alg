package main

import "fmt"

//给定一个二叉树，它的每个结点都存放着一个整数值。
//
//找出路径和等于给定数值的路径总数。
//
//路径不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
//
//二叉树不超过1000个节点，且节点数值范围是 [-1000000,1000000] 的整数。
//
//示例：
//
//root = [10,5,-3,3,2,null,11,3,-2,null,1], sum = 8
//
//10
///  \
//5   -3
/// \    \
//3   2   11
/// \   \
//3  -2   1
//
//返回 3。和等于 8 的路径有:
//
//1.  5 -> 3
//2.  5 -> 2 -> 1
//3.  -3 -> 11

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode,sum int)int{
	if root == nil {
		return 0
	}
	return pathFrom(root,sum)+pathFrom(root.Left,sum)+pathFrom(root.Right,sum)
}

func pathFrom(root *TreeNode,sum int)int{
	if root == nil {
		return 0
	}
	count := 0

	if root.Val == sum {
		count++
	}

	//重复子问题
	return count+pathFrom(root.Left,sum-root.Val)+pathFrom(root.Right,sum-root.Val)
}



func main(){
	root := &TreeNode{Val: 10}
	left := &TreeNode{Val: 5}
	right := &TreeNode{Val: -3}
	root.Left = left
	root.Right = right

	left_left := &TreeNode{Val: 3}
	left_right := &TreeNode{Val: 2}
	right_right := &TreeNode{Val: 11}
	root.Left.Left = left_left
	root.Left.Right = left_right
	root.Right.Right = right_right

	left_left.Left = &TreeNode{Val: 3}
	left_left.Right = &TreeNode{Val:-2}
	left_right.Right = &TreeNode{Val: 1}

	fmt.Println(pathSum(root,8))
}



