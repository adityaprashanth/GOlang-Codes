package main

import "fmt" 
       
func main() {

              var x int = 5

              fmt.Println(x)

              fmt.Println(f1(&x)) // pass by reference

              fmt.Println(x)
}


func f1(x *int) int { *x++
                      return *x }
