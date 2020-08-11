package main

import "fmt"

func rightBound(nums []int,target int)int{
	left,right := 0,len(nums)-1
	for left <= right{
		mid := left + (right-left)/2
		if nums[mid] == target {
			left = mid+1
		}else if nums[mid] < target {
			left = mid + 1
		}else if nums[mid] > target {
			right = mid -1
		}
	}
	//left = right+1
	if right >= len(nums) || nums[right] != target {
		return -1
	}

	return right
}


func main(){
	nums := []int{1,2,2,2,3}
	fmt.Println(rightBound(nums,2))
}
