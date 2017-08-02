package main

import (
	"fmt"
)


func apply( nums []int, f func (int) int) func() {
	
	for i, n := range nums{
		nums[i] = f(n)
	}
	
	return func(){
		
		fmt.Println(nums)
	}
	
}



func main() {
	nums := []int{1,2,3,4}
	
	result :=apply( nums, func ( a int) int { 
			return a*2
		})
	result()
}
