package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456"

func main() {
	fmt.Println("Welcome to handling url in golang")

	//parsing url
	result, err := url.Parse(myurl)

	if err != nil {
		panic(err)
	}

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port()) 
	// fmt.Println(result.RawQuery)

	qparams := result.Query()

	fmt.Printf("type of query parameters %T \n", qparams)
	fmt.Println(qparams["coursename"])

	for _, value := range qparams {
		fmt.Println(value)
	}

	// making url
	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user-hitesh",
	}

	anotherUrl := partsOfUrl.String()

	fmt.Println(anotherUrl)
}