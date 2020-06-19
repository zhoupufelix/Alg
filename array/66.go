package main

import "fmt"

//给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
//最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。你可以假设除了整数 0 之外，这个整数不会以零开头。
//输入: [1,2,3]
//输出: [1,2,4]
//解释: 输入数组表示数字 123。

func plusOne(arr []int) (res bool){
	for i:=len(arr)-1 ;i>=0;i--{
		if arr[i] == 9 {
			arr[i] = 0
		}else{
			arr[i]++
			break
		}
	}
	if arr[0] == 0 {
		res = true
	}
	return res
}


func main(){
	tmp := []int{1}
	arr := []int{1,2,3,9}
	res := plusOne(arr)
	if res {
		tmp = append(tmp,arr...)
	}else{
		tmp = arr
	}
	fmt.Println(tmp)
}
