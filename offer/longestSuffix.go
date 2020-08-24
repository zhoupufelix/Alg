package main

import (
	"fmt"
	"strings"
)

//题目14: 最长公共前缀
//编写一个函数来查找字符串数组中的最长公共前缀。如果不存在公共前缀，则返回""
//示例1:
//
//输入: ["flower","flow","flight"]
//输出: "fl"
//示例 2:
//
//输入: ["dog","racecar","car"]
//输出: ""
//解释: 输入不存在公共前缀。
//解释:
//
//输入不存在公共前缀。
//说明：
//
//所有输入只包含小写字母 a-z

func longestSuffix(strs []string)string{
	if len(strs) < 1 {
		return ""
	}
	prefix := strs[0]
	for _,str := range strs{
		for strings.Index(str,prefix) != 0 {
			if len(prefix) == 0 {
				return ""
			}
			//压缩一位继续判断
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}


func main(){
	strs := []string{"flow","flower","flight"}
	fmt.Println(longestSuffix(strs))
}
