package main

import "fmt"

//在一个二维数组中，每一行都按照从左到右递增的顺序排序，
//每一列都按照从上到下递增的顺序排序。请完成一个函数，输入这样的一个
//二维数组和一个整数，判断数组中是否含有该整数。
//1 2 8 9
//2 4 9 12
//4 7 10 13
//6 8 11 15

func Find(matrix [][]int,target int)bool{
	for i,j := len(matrix)-1,0;i<len(matrix) && j>=0;{
		if  matrix[i][j] == target {
			return true
		}else if matrix[i][j] > target{
			i--
		}else if matrix[i][j] < target {
			j++
		}
	}
	return false
}



func main(){
	martrix := [][]int{
		{1, 2, 8, 9},
		{2, 4, 9, 12},
		{4, 7, 10, 13},
		{6, 8, 11, 15},
	}
	fmt.Println(Find(martrix,12))

}

