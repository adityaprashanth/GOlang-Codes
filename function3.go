package main

import "fmt" 
       
func main() {

              var x int = 5

              r1, r2 := f1(&x) // pass by reference

              fmt.Println(r1, r2)
}


func f1(x *int) (int, int) { r1 := *x + 1
                             r2 := *x - 1
                             return r1, r2 }

