package main

import (
	"fmt"
)

var result [][]int

func subset(nums []int){
	//路径
	track := []int{}
	backtrack(nums,0,track)
}

func backtrack(nums []int,start int,track []int){
	subRes := make([]int, len(track))
	copy(subRes, track)

	//走过的路径就是子集
	result = append(result,subRes)
	//选择列表
	for i:=start;i<len(nums);i++{
		//做选择
		track = append(track,nums[i])
		//回溯
		backtrack(nums,i+1,track)
		//撤销选择
		track   = track[:len(track)-1]
	}
}


func main(){
	nums := []int{1,2,3}
	subset(nums)
	fmt.Println(result)
}