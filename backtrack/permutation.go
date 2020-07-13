package main

import "fmt"

var result [][]int


//nums = [1,2,3] 求数组的全排列
func permute(nums []int){
	//路径
	track := []int{}
	//选择
	used  := make([]bool,len(nums))
	//回溯
	backtrack(nums,track,used)
}

func backtrack(nums []int,track []int,used []bool){
	//到达叶子节点，长度等于数组长度
	if len(track) == len(nums) {
		tmp := make([]int,len(nums))
		copy(tmp,track)
		result = append(result,tmp)
		return
	}

	for i:=0;i<len(nums);i++{
		if used[i] == true {
			continue
		}
		//做选择
		used[i] = true // 标记已使用
		track = append(track,nums[i])
		//回溯
		backtrack(nums,track,used)
		//撤销选择
		track = track[:len(track)-1]
		used[i] = false
	}
}


func main(){
	nums := []int{1,2,3}
	permute(nums)
	fmt.Println(result)
}