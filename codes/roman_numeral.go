package codes

import (
	//"fmt"
)

func RomanNumeral(s string)  int {
	s="CMLII"
	reflect:=make(map[string]int,0)
	reflect["M"]=1000
	reflect["D"]=500
	reflect["C"]=100
	reflect["L"]=50
	reflect["X"]=10
	reflect["V"]=5
	reflect["I"]=1
	res :=[]byte(s)
	before :=res[0]
	result :=0
	for i:=range res{
		result += reflect[string(res[i])]
		if before!=res[i]&&reflect[string(before)]<reflect[string(res[i])]{
			result-=reflect[string(before)]*2
		}
		before=res[i]
	}
	return result
}