package main

import "fmt"

//给定一个无序的整数数组，找到其中最长上升子序列长度
//[10,9,2,5,3,7,101,18]，它的长度是4

func lengthOfLIS(nums []int)int{
	n := len(nums)
	if n <= 0 {
		return 0
	}
	//初始化dptable 默认是1 就是其本身
	dp := make([]int,n)
	for k:=0;k<n;k++{
		dp[k] = 1
	}


	for i:=0;i<n;i++{
		for j:=0;j<i;j++{
			if  nums[i] > nums[j]{
				dp[i] = max(dp[i],dp[j]+1)
			}
		}
	}

	lis := 0
	for m:=0;m<n;m++{
		lis = max(lis,dp[m])
	}
	return lis
}




func max(a,b int)int{
	if a < b {
		return b
	}
	return a
}



func main(){
	nums := []int{10,9,2,5,3,7,101,18}
	fmt.Println(lengthOfLIS(nums))
}



