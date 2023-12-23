package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app: ")
	fmt.Println("please rate our pizza from 1 - 5: ")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("thanks for rating:", input)
	// converting string to int now if we didnt use trim space we had a problem that
	// when we pressed enter so that was also considered in our input so we had to trim that
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("increased rating: ", numRating + 1)
	}
}