package main

import "fmt"

func main() {
	fmt.Println("Slices: ")

	var fruit = []string{} //slices
	fmt.Printf("type %T\n", fruit)

	// initialising slices
	var list =[]string{"mango", "banana", "guava"}

	// adding new items
	list = append(list, "orange", "apple")
	fmt.Println(list)

	// slicing the slices
	list = append(list[1:])
	fmt.Println(list)

	//using memory management make() new() for initialising slices
	highscores := make([]int, 4)

	highscores[0] = 234
	highscores[1] = 232
	highscores[3] = 321

	fmt.Println(highscores)

}