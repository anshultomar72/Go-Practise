package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Reading and writing from file using go")
	content := "I am Anshul Tomar."

	file, err := os.Create("./test.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println(length)

	defer file.Close()

	readFile("./test.txt")
}

func readFile(fileName string) {
	dataByte, err := os.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(dataByte))

}
