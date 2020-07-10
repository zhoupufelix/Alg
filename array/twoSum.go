package main

import (
	"fmt"
	"sort"
)

//如果假设输入一个数组 nums 和一个目标和 target，请你返回 nums 中能够凑出 target 的两个元素的值，比如输入 nums = [5,3,1,6], target = 9，那么算法返回两个元素 [3,6]。
//可以假设只有且仅有一对儿元素可以凑出 target。

func twoSum(nums []int,target int)[][]int{
	var result [][]int
	//排序，实现比较方法即可
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	sum,l,r := 0,0,len(nums)-1

	for l < r {
		left,right := nums[l],nums[r]
		sum = left+right
		if sum < target {
			for  l < r && nums[l] == left  { l++ }
		}else if  sum > target{
			for l < r && nums[r] == right { r-- }
		}else if sum == target {
			result = append(result,[]int{nums[l],nums[r]})
			for  l < r && nums[l] == left  { l++ }
			for l < r && nums[r] == right { r-- }
		}
	}
	return result
}


func main(){
	nums := []int{1,1,1,2,2,3,3}
	fmt.Println(twoSum(nums,4))
}
