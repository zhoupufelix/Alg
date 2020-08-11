package main

import "fmt"

func binarysearch(nums []int,target int)int{
	left,right := 0,len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if  nums[mid] == target{
			return mid
		}else if  nums[mid] < target{
			left = mid+1
		}else if nums[mid] > target {
			right = mid-1
		}
	}
	return -1
}

func main(){
	nums := []int{1,2,2,2,3}
	fmt.Println(binarysearch(nums,2))
}
