package codes

import (
)

func TowSum(nums []int,target int) []int {
	if len(nums)<=1{
		return []int{0,0}
	}
	var result =make(map[int]int,0)
	for i :=range nums{
		result[nums[i]]=i
	}

	for j:=range nums {
		tmp :=target-nums[j]
		if _,ok:=result[tmp];ok{
			if result[tmp]!=j {
				return []int{j, result[tmp]}
			}
		}
	}
	return []int{}
}
