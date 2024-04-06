package main

import (
        "fmt" 
        "time"
       )

func main() {
 
     type Employee struct {
                            ID int
                            Name string
                            Address string
                            DoB time.Time
                            Position string
                            Salary int
                            ManagerID int
                          }

     var dilbert, david Employee

     dilbert.ID = 1
     dilbert.Name = "Dilbert"
     dilbert.Salary = 10000
    
     fmt.Println(dilbert)

     dilbert.Salary += 1000

     fmt.Println(dilbert)

     p := &dilbert.Position

     *p = "Manager"

     fmt.Println(dilbert)

     fmt.Println(dilbert.Salary == david.Salary)

}
