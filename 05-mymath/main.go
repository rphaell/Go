package main

import (
	"fmt"
	"math/big"

	// "math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("Welcome to maths ingolang")
	// var mynumberOne int = 2
	// var mynumbertwo float64 = 4

	// fmt.Println("The sum is: ", mynumberOne+int(mynumbertwo))

	// random number
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(6) + 1)

	// random from crypto
	myRandomNumber, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNumber)
}
