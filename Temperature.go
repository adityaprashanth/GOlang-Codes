package main // all prg will have it

import "fmt" 

func main() {

	const freez = 32.0
	const boil = 212.0

	fmt.Printf("%g = %g\n",freez,fToc(freez))
	fmt.Printf("%g = %g\n",boil,fToc(boil))
	
}

func fToc(f float64)float64 {

	return 5.0/9.0 * (f-32)

}