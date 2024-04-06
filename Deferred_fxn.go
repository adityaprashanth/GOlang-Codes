package main

import ("fmt"
	 "time"
	"os"
)

func main() {
	defer fmt.Println("good ")
	defer fmt.Println("morning ")

	fmt.Println("hello ")
}

//output :  hello good morning