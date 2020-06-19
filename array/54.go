package main

import "fmt"

//给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。
//在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。
//找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

//输入：[1,8,6,2,5,4,8,3,7] 输出：49

func MaxArea(arr []int)int{
	//i 左边  j 右边 res 面积
	i,j,res := 0,len(arr)-1,0
	for i<j {
		if  arr[i] < arr[j]{
			res = Max(res,(j-i)*arr[i])
			i++
		}else{
			res = Max(res,(j-i)*arr[j])
			j--
		}
	}
	return res
}


func Max(a,b int)int{
	if a > b {
		return a
	}
	return b
}

func main(){
	arr := []int{1,8,6,2,5,4,8,3,7}
	fmt.Println(MaxArea(arr))
}