package main

import "fmt"

func main() {
	fmt.Println("welcome to my functions")
	greeter()

	result := adder(3, 9)
	fmt.Println("result is: ", result)

	proResult := proAdder(2, 5, 7, 4)
	fmt.Println("pro result is: ", proResult)

}

func greeter() {
	fmt.Println("Whatsupp")
}

func adder(value1 int, value2 int) int {
	return value1 + value2

}

func proAdder(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}

	return total
}

func greetertwo() {
	fmt.Println("eae")
}
