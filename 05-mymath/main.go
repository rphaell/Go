package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to maths ingolang")
	// var mynumberOne int = 2
	// var mynumbertwo float64 = 4

	// fmt.Println("The sum is: ", mynumberOne+int(mynumbertwo))

	// criando numero random de 1 a 6
	max := 6
	min := 1
	rand.Seed(time.Now().UnixNano())
	v := rand.Intn(max-min) + min
	fmt.Println(v)

	//ou...

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is :", diceNumber)

	// random from crypto
	// myRandomNumber, _ := rand.Int(rand.Reader, big.NewInt(5))
	// fmt.Println(myRandomNumber)
}
