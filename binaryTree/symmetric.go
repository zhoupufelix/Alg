package main


import "fmt"




type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}




//给定一个二叉树，检查它是否是镜像对称的。
//例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
func isSymmetric(root *TreeNode)bool{
	if root == nil  {
		return false
	}
	return isMirrorRecursive(root.Left,root.Right)
}





func isMirrorRecursive(left,right *TreeNode)bool{
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	//重复子问题
	//两个根节点是否相等，a 的左节点的值是否等于右节点的值
	return left.Val == right.Val && isMirrorRecursive(left.Left,right.Right) && isMirrorRecursive(left.Right,right.Left)
}


func main(){
	root := &TreeNode{Val: 1}
	left := &TreeNode{Val: 2}
	right := &TreeNode{Val: 2}
	root.Left = left
	root.Right = right
	left_left := &TreeNode{Val: 3}
	left_right := &TreeNode{Val: 4}
	right_left := &TreeNode{Val: 4}
	right_right := &TreeNode{Val: 3}
	root.Left.Left = left_left
	root.Left.Right = left_right
	root.Right.Left = right_left
	root.Right.Right = right_right

	fmt.Println(isSymmetric(root))
}