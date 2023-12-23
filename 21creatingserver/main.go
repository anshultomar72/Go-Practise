package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// fmt.Println("Creating webserver")
	// PerformGetRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myUrl = "http://localhost:3000/get"

	response, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("status code is: ", response.StatusCode)

	// fmt.Println("response length is: ", response.ContentLength)
	// content, _ := io.ReadAll(response.Body)
	// fmt.Println(string(content))

	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println(byteCount)
	fmt.Println(responseString.String())
}

func PerformPostJsonRequest(){
	const myUrl = "http://localhost:3000/post" 

	//fake json payload
	requestBody := strings.NewReader(`
		{
			"coursename" : "Lets go with golang",
			"price" : 0,
			"platform" : "learnCodeOnline.in"
		}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myUrl = "http://localhost:3000/postform" 

	//fake formdata

	data := make(url.Values)
	data.Add("firstname", "hitesh")
	data.Add("lastname", "choudhary")
	data.Add("email", "hitesh@go.dev")

	response, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}