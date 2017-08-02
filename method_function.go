package main

import (
	"fmt"
)

type name struct {
	first, last string
}

func (n name ) show() {
	fmt.Println( n.first , n.last)
}

func main() {
	gal := name {
		first : "hemant",
		last : "sachdeva",
	}
	gal.show()
}
