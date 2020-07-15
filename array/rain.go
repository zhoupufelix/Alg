package main

import "fmt"

//接雨水
//给定n个非负整数表示每个宽度为1的柱子的高度图，计算按此排列的柱子，
//下雨之后能接多少雨水。
//示例：
//输入 ：[0,1,0,2,1,0,1,3,2,1,2,1]
//输出 ：6

func trap(nums []int)int{
	l,r:= 0,len(nums)-1
	lmax,rmax,ans := 0,nums[r],0

	//双指针
	for l<=r {
		//左边最高
		lmax = max(lmax,nums[l])
		//右边最高
		rmax = max(rmax,nums[r])

		if lmax < rmax {
			ans += lmax - nums[l]
			l++
		}else{
			ans += rmax - nums[r]
			r--
		}
	}
	return ans
}


func max(a,b int)int{
	if a > b  {
		return a
	}
	return b
}



func main(){
	nums := []int{0,1,0,2,1,0,1,3,2,1,2,1}
	fmt.Println(trap(nums))
}