package codes

import "strconv"

func isPalindrome(x int) bool  {

	var str = strconv.Itoa(x)
	strB :=[]byte(str)

	for i :=range strB{
		if strB[i]!=strB[len(strB)-i-1]{
			return false
		}

	}
	if x<0 {
		return false
	}

	return true
}