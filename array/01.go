package main

import "fmt"

//给定一个整数数组nums 和一个目标值 target ,请你在该数组中找出和目标值的那两个整数，
//并返回他们数组下标，你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
//示例：
//给定 nums = [2,7,11,15],target = 9
//因为 nums[0] + nums[1] = 2+7 = 9
//所以返回[0,1]

func twoSum(nums []int,target int)[]int{
	m:=make(map[int]int)
	var result []int
	for i:=0;i<len(nums);i++{
		if v,exist := m[target-nums[i]];exist{
			result = append(result,i,v)
		}else{
			m[nums[i]] = i
		}
	}
	return result
}


func main(){
	nums := []int{2,7,11,15}
	fmt.Println(twoSum(nums,9))
}
