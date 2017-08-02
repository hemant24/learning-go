package main

import (
	"fmt"
)

func div(op0, op1 int) (q, r int) {
	r = op0
	for r >= op1 {
		q++
		r = r - op1;	
	}
	return
}

func main() {
	q, r := div(52,5)
	
	fmt.Printf("div(52,5) %d \n", q)

	fmt.Printf("div(52,5) %d \n", r)
}
