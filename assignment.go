package main

import "fmt"

func main() {

     var i, j, k = 2, 3, 4 // tuple assignment

     fmt.Println(i, j, k)

     i, j = j, i // swapping the values

     fmt.Println(i, j, k)

     medals := []string{"gold", "silver", "bronze"}

     fmt.Println(medals[0], medals[1], medals[2])

}
