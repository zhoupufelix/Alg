package main

import "fmt"

//输入: nums1 = [1,2,2,1], nums2 = [2,2]
//
//输出: [2,2]
//示例 2:
//
//输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//
//输出: [4,9]


//进阶:
//
//如果给定的数组已经排好序呢？将如何优化你的算法呢？

func intersect(nums1,nums2 []int)[]int{
	m := map[int]int{}
	for _,v := range nums1 {
		m[v]++
	}
	k:=0
	for _,v := range nums2{
		if _,ok :=m[v];ok {
			m[v]--
			nums2[k] = v
			k++
		}
	}
	return nums2[:k]
}

//如果是拍好序的数组，使用双指针p,q 指向两个数组头部 ，
//如果指针指向的数不相等，移动较小的指针，如果一样，两个都移动


func main(){
	nums1 := []int{1,2,2,1}
	nums2 := []int{2,2}
	fmt.Println(intersect(nums1,nums2))
}
