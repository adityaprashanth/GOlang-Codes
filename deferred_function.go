package main

import "fmt"

func main() {

	defer fmt.Println("Morning")
	defer fmt.Println("Good")

	fmt.Println("hello")
}
