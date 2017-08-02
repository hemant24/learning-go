package main

import (
	"fmt"
)

type matrix [2][2]byte




func main() {
	var mat1 matrix
	mat1 = initMat()
	fmt.Println(mat1)
}

func initMat() matrix {
	
	return matrix {
		{1,2} , {3,4},
	}
	

}
