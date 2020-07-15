package main

import "fmt"

//你是一个专业的盗贼，计划偷打劫街边的房屋。每间房屋都藏有一定的现金，
//影响你的唯一制约因素就是相邻的房屋有相互连通的防盗系统，如果两间相邻的
//房屋在同一晚上被盗贼闯入，系统会自动报警。
//给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，
//能够偷窃到的最高金额。
//示例：
//输入：[1,2,3,1]
//输出：4
//解释：偷窃 1 号房屋（金额 = 1），然后偷窃3号房屋（金额 = 3）。
//偷窃到的最高金额 = 1 + 3 = 4

var memo  = map[int]int{}

func rob(nums []int)int{
	//备忘录
	return dp(nums,0)
}

func dp(nums []int,start int)int{
	//base case
	if start >= len(nums) {
		return 0
	}
	if v,ok := memo[start] ; ok  {
		return v
	}
	//当前状态 抢或不抢
	sum := max(
		// 不抢，去下家
		dp(nums, start + 1),
		// 抢，去下下家
		nums[start] + dp(nums, start + 2) )

	//记录备忘录
	memo[start] = sum
	return sum
}


func max(a,b int)int{
	if a > b {
		return a
	}
	return b
}


func main(){

	nums := []int{2,7,9,3,1}
	fmt.Println(rob(nums))
}

