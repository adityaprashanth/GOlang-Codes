package main

import "fmt" 
       
func main() {
    fmt.Println(divide(10, 2))

    fmt.Println(divide(10, 0))
    fmt.Println("This line will not execute")
}


func divide(a, b int) int {

    if b == 0 {
        panic("Divide by zero")        
        if r := recover(); r != nil {
            fmt.Println("In recover ", r)
        }     

    }

    return a / b
}
