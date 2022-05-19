package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Thanks for buying from us"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our store:")

	// comma ok // error err

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Rating type is: %T", input)

}
