package main

import "fmt"

func main() {
	fmt.Println("Welcome")

	// var ptr *int

	var number = 5;
	ptr := &number

	fmt.Println(ptr)
	fmt.Println(*ptr)

	*ptr = *ptr * 2

	fmt.Println(*ptr)
}