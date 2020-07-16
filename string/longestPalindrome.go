package main

import "fmt"

func palindrome(s string,l,r int)string{
	str := []byte(s)
	length := len(s)
	for l>=0 && r < length && str[l]==str[r]{
		l--
		r++
	}
	str = str[l+1:r]
	s = string(str)
	return s
}

func longestPalindrome(s string)string{
	var res string
	for i:=0;i<len(s);i++{
		s1 := palindrome(s,i,i)
		s2 := palindrome(s,i,i+1)

		if len(res) < len(s1) {
			res = s1
		}

		if len(res) < len(s2) {
			res = s2
		}

	}
	return res
}





func main(){
	s := "babadab"
	fmt.Println(longestPalindrome(s))
}