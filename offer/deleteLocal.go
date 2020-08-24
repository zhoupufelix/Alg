package main

import "fmt"

//题目27：移除元素
//给定一个数组 nums 和一个值 val，你需要原地移除所有数值等于 val 的元素，返回移除后数组的新长度。
//不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
//
//元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
//
//示例 1:
//
//给定 nums = [3,2,2,3], val = 3,
//函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。
//你不需要考虑数组中超出新长度后面的元素。



//类似题目
//题目26：删除排序数组中的重复项
//给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次。
//返回移除后数组的新长度。不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
//
//示例 1:
//
//给定数组 nums = [1,1,2],
//函数应该返回新的长度 2, 并且原数组 nums 的前两个元素被修改为 1, 2。
//你不需要考虑数组中超出新长度后面的元素。
//示例 2:
//
//给定 nums = [0,0,1,1,1,2,2,3,3,4],
//函数应该返回新的长度 5, 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4。
//你不需要考虑数组中超出新长度后面的元素。




func removeElement(nums []int,val int)int{
	length := len(nums)
	for i:= 0;i<length{
		if nums[i] == val {
			nums = append(nums[:i],nums[i+1:]...)
		}else{
			i++
		}
	}
	return len(nums)
}


func removeDuplicate(nums []int)int{
	for i:=0;i<len(nums)-1{
		if nums[i+1] == nums[i] {
			nums = append(nums[:i],nums[i+1:]...)
		}else{
			i++
		}
	}
	return len(nums)
}


func main(){
	nums := []int{3,2,2,3}
	fmt.Println(removeElement(nums,3))
	nums2 := []int{0,0,1,1,1,2,2,3,3,4}
	fmt.Println(removeDuplicate(nums2))
}