package main

import (
	"fmt"
	"strings"
)


func cap( p *struct { first, last string}){

	p.first = strings.ToUpper(p.first)
	p.last = strings.ToUpper(p.last)
}


func main() {

	
	p := &struct{first,last string} { "hemant", "sachdeva"}
	
	
	cap(p)
	
	fmt.Println(*p)
}

