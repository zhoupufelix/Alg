package main

import (
	"fmt"
)

//给定一个数组，它的第i个元素是一支给定的股票在第i天的价格。
//设计一个算法来计算你所能获取的最大利润。你最多可以完成一笔交易。
//注意：你不能同时参与躲避交易（你必须在再次购买前出售掉之前的股票）
//k = 1

func maxProfit(prices []int)int{
	n  := len(prices)
	//初始化dp table
	dp := make([][]int,n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}


	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i:=1;i<n;i++{
		//不持有股票 前一天就不持有（保持），前一天持有，今日卖出
		dp[i][0] = max(dp[i-1][0],dp[i-1][1]+prices[i])
		//持有股票 前一天就持有（保持），前一天没持有，今日买入
		dp[i][1] = max(dp[i-1][1],-prices[i])
	}
	fmt.Println(dp)
	return dp[n-1][0]
}

func max(a,b int)int{
	if a>b {
		return a
	}
	return b
}

func main(){
	prices := []int{7,1,5,3,6,4}
	fmt.Println(maxProfit(prices))
}