package main

import "fmt"

func main() {
	fmt.Println("Arrays")

	var fruitlist [4]string

	fruitlist[0] = "anshul"
	fruitlist[1] = "tomar"
	fruitlist[3] = "yes"

	fmt.Println(fruitlist)
	fmt.Println(len(fruitlist))

	var veglist = [3]string{"mushroom", "bhindi", "cabbage"}
	fmt.Println(veglist)


}