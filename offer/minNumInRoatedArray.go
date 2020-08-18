package main

import "fmt"

//把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
//输入一个非减排序的数组的一个旋转，输出旋转数组的最小元素。
//例如数组 {3,4,5,1,2} 为 {1,2,3,4,5} 的一个旋转，该数组的最小值为 1。
//
//**NOTE：**给出的所有元素都大于 0，若数组大小为 0，请返回 0。
func Solution(nums []int)int{
	l := len(nums)
	if l == 0 {
		return 0
	}
	p,q := 0,l-1
	mid := p
	for nums[p] >= nums[q]{
		//如果距离只相差1 因为是非减序列 p就是最小的元素
		if q-p == 1 {
			mid = p
			break
		}
		
		mid = p+(q-p)/2
		if nums[p] == nums[q] && nums[mid] == nums[p] {
			mid = getMinIndex(nums,p,q)
			break
		}
		//证明mid在左边的递增序列中 最小值在mid右侧 压缩区间
		if nums[mid] >= nums[p]{
			p = mid
		}else if nums[mid] <= nums[q] {
			q = mid
		}
	}
	return nums[mid]
}

func getMinIndex(nums []int,p,q int)int{
	min := 0
	for i:=p+1 ;i<=q;i++{
		if nums[i] < min {
			min = i
		}
	}
	return min
}


func main(){
	nums := []int{3,4,5,1,2}
	fmt.Println(Solution(nums))
}