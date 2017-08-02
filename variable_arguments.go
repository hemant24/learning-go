package main

import (
	"fmt"
)

func avg(nums ...float64) float64 {
	n := len(nums)
	t := 0.0
	for _, v := range nums {
		t += v
	}
	return t / float64(n)

}

func main() {
	fmt.Printf("avg %f\n", avg(1, 2, 3, 4))

	points := []float64{1, 2, 3, 4}

	fmt.Printf("avg %f\n", avg(points...))
}
