package main

import (
	"fmt"
	"time"
        "strconv"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("At %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {

        i, err := strconv.Atoi("42A")
        if err != nil {
                        fmt.Printf("couldn't convert number: %v\n", err)
                        return
                      }
        fmt.Println("Converted integer:", i)

	if err = run(); err != nil {
		fmt.Println(err)
	}

}