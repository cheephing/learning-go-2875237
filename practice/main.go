package main

import (
	"fmt"
)

const aConst int = 64

func main() {

	var aString string = "Hello, World! rrr"
	fmt.Println(aString)
	fmt.Printf("Type: %T\n", aString)

	var aInterger int = 42
	fmt.Println(aInterger)

	var defaultInt int
	fmt.Println(defaultInt)

	var anotherString ="this is a string"
	fmt.Println(anotherString)
	fmt.Printf("Type: %T\n", anotherString)

	var anotherInterger = 4
	fmt.Println(anotherInterger)
	fmt.Printf("Type: %T\n", anotherInterger)

	mystring := "this is a string"
	fmt.Println(mystring)
	fmt.Printf("Type: %T\n", mystring)

	fmt.Println(aConst)
	fmt.Printf("Type: %T\n", aConst)
}
