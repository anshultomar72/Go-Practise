package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study")

	presentTime := time.Now()

	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("02-01-2006 15:04::05 Monday"))

	createdDate := time.Date(2020, time.August, 10, 23, 23, 0, 0, time.UTC)

	fmt.Println(createdDate)
}