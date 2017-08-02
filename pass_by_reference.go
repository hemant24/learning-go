package main

import (
	"fmt"
	"math"
)

func double(val *float64) {
	*val = *val * 2
	fmt.Printf("val is %.5f \n", *val)
	
}

func main() {
	pi := math.Pi
	
	fmt.Printf("before dbl() p = %.5f\n", pi)
	
	double(&pi)
	
	fmt.Printf("after dbl() p = %.5f\n", pi)
}
