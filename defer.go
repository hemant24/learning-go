package main

import (
	"fmt"
)


func show(names ...string)  {
	defer fmt.Println("All Done")
	
	for _, name := range names {
	
		fmt.Println(name)
	}
	
	
}



func main() {
	names := []string{"hemant", "parul"}
	
	show(names...)
}
