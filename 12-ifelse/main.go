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

	if 10%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 20; num < 10 {
		fmt.Println("number is less than 10")
	} else {
		fmt.Println("number is greater than 10")
	}

	// if err != nil {

	// }

}
