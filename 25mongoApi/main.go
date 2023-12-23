package main

import (
	"25mongoApi/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Mongo API")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":3000", r))
	fmt.Println("Listening on port 3000...")
}
