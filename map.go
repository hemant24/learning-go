package main

import (
	"fmt"
)

func main() {
	var nameMap map[int]string
	
	nameMap = make(map[int]string)
	nameMap[1]= "hemant"
	nameMap[2]="Rohit"
	fmt.Println(nameMap)
}
