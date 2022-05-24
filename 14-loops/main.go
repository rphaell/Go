package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops!")

	days := []string{"Sunday", "Monday", "Tuesday", "Friday"}

	fmt.Println(days)

	// TYPE LOOP 1
	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for _, day := range days {
		fmt.Printf("index is and value is %v\n", day)
	}

	rougueValue := 1

	for rougueValue < 10 {

		if rougueValue == 5 {
			rougueValue++
			continue
		}

		fmt.Println("Value is : ", rougueValue)
		rougueValue++
	}

}
