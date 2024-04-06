package main // all prg will have it

 import "fmt"

 func main() {

	x:=1
	y:=0

	p := &x  		// p is pointing to address of x whose value is 1

	fmt.Println(*p)

	*p = 2 			// the value p is pointing to is changed to 2 

	fmt.Println(x)

	q := &x   		// q is pointing to address of x whose is in turn being pointed by p

	fmt.Println(p == q)

	q = &y

	fmt.Println(*q)

	*q = 7

	fmt.Println(x)	

	l := new(int) 		// l is pointing to int whose default is 0

	fmt.Println(*l)

	*l = 2			// the value l is pointing to is changed to 2 

	l1 := *l

	fmt.Println(l1)	

}

/*

func incr(p *int) int {

     *p++ // increments what p points to

     return *p
}
*/
