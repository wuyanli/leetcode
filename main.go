package main

import (
	"./codes"
	"fmt"
)

func main() {
	 nums := []int{3,3,3}
	 result := make([]int,0)
	 result=codes.TowSum(nums,6)
	 fmt.Println(result)
}
