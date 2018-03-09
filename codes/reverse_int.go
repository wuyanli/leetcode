package codes

import (
	"strconv"
	"math"
)

func Reverse(num int) int  {
	if num>math.MaxInt32||num<math.MinInt32{
		return 0
	}
	flag :=false
	if num<0 {
		num=num*-1
		flag=true
	}
	res :=strconv.Itoa(num)
	resB :=[]byte(res)
	resC :=make([]byte,0)
	for i:=len(resB)-1;i>=0;i--{
		resC=append(resC, resB[i])
	}
	res = string(resC)
	result :=0
	result,_=strconv.Atoi(res)
	if flag{
		result=result*-1
	}
	if result>math.MaxInt32||result<math.MinInt32{
		return 0
	}
	return result
}