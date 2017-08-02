package main

import (
	"fmt"
	"strings"
	"errors"
)

func main() {
	name, error := toUpper("");
	if error != nil{
		fmt.Println("error is : " , error)
	}else{
		fmt.Println(name)
	}
}


func toUpper(name string) (string, error) {
	if name == "" {
		return "", errors.New("Invalid name")
	}
	return strings.ToUpper(name) , nil

}