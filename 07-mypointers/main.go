package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers")

	// criando um pointer sem valor
	// var ptr *int
	// fmt.Println("value of pointer is:", ptr)

	myNumber := 23

	// criando um pointer referenciando a variavel myNumber
	var ptr = &myNumber
	fmt.Println("value of actual pointer is: ", ptr)
	fmt.Println("value of actual pointer is: ", *ptr)

	// multiplicando um pointer criado por 2
	*ptr = *ptr * 2
	fmt.Println("new values is: ", myNumber)
}
