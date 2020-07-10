package main

import (
	"fmt"
	"sort"
)

func twoSum(nums []int,start,target int)(result [][]int){
	//先排序
	sort.Slice(nums, func(i, j int) bool {
		return nums[i]<nums[j]
	})

	l,r := start,len(nums)-1
	for l<r {
		left,right := nums[l],nums[r]
		sum := left + right
		if sum < target {
			l++
		}else if sum > target  {
			r--
		}else if sum == target {
			result = append(result,[]int{left,right})
			for l<r && nums[l] == left{l++}
			for l<r && nums[r] == right{r--}
		}
	}

	return result
}


//三数之和，先选定一个数
func threeSum(nums []int,target int)(result [][]int){
	//先排序
	sort.Slice(nums, func(i, j int) bool {
		return nums[i]<nums[j]
	})

	length := len(nums)

	for i:=0;i<length;i++ {
		tuple := twoSum(nums,i+1,target-nums[i])
		if len(tuple) > 0 {
			for j:=0;j<len(tuple);j++{
				tuple[j] = append(tuple[j],nums[i])
				result = append(result,tuple[j])
			}

		}
		//跳过重复的数
		for i < length - 1 && nums[i] == nums[i + 1] {i++}
	}

	return result
}


func main(){
	nums := []int{1,1,1,2,3,4}
	fmt.Println(threeSum(nums,6))
}
