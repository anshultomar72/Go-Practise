package main

import "fmt"

func main() {
	fmt.Println("Removing slices")

	var courses = []string{"javascript", "go", "ruby", "python", "cpp"}
	fmt.Println(courses)

	var index = 2
	courses = append(courses[:index], courses[index+1:]...) //dots to include more arguments
	fmt.Println(courses)
}