package main

import "fmt"

func leftBound(nums []int,target int)int{
	left,right := 0,len(nums)-1
	for left <= right{
		mid := left+(right-left)/2
		if  nums[mid] == target{
			right = mid-1
		}else if nums[mid] > target {
			right = mid-1
		}else if  nums[mid] < target{
			left = mid+1
		}
	}
	if left >= len(nums) || nums[left] != target {
		return -1
	}
	
	return left
}


func main(){
	nums := []int{1,2,2,2,3}
	fmt.Println(leftBound(nums,2))
}