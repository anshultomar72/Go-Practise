package main

import (
	"fmt"
	"io"
	"net/http"
)
const url = "https://lco.dev"

func main() {
	fmt.Println("Web requests ")

	resp, err := http.Get(url)
	
	if err != nil {
		panic(err)
	}
	fmt.Printf("%T \n", resp)

	defer resp.Body.Close() // caller's responsibility to close the connection

	databytes, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}
	
	fmt.Printf("%T \n", databytes)
	fmt.Println(string(databytes))
}