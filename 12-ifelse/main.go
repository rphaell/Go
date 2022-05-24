package main

import "fmt"

func main() {
	fmt.Println("welcome to if and else in go")

	LoginCount := 10
	var result string

	if LoginCount < 10 {
		result = "Regular User"
	} else if LoginCount > 10 {
		result = "Heavy User"
	} else {
		result = "exactly 10 login"
	}

	fmt.Println(result)

	numero := 4
	if numero%2 == 0 {
		result = "par"
	} else {
		result = "impar"
	}
	fmt.Println(result)

	if num := 20; num < 10 {
		fmt.Println("nota ruim")
	} else {
		fmt.Println("nota boa")
	}

	// if err != nil {

	// }

}
