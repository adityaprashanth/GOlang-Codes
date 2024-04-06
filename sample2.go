package main // all prg will have it

 import "fmt"

 func main() {

	fmt.Println("hello") 


	var i int		// gives 0 
	var i1 = 32
	
	var s string		// gives 1 whitespace 
	var s1 = "hello"
	
	var f float32		// gives 0.000000  also for float give 32/64
	var f1 = 3.14

	fmt.Printf("x = %d , x1 = %d\n",i,i1) 
	fmt.Printf("y = %s , y1 = %s\n",s,s1)
	fmt.Printf("y = %f , y1 = %f",f,f1) 
}
