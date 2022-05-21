package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("switch and case in go!")

	// código para o dado de 1 a 6 com precisão
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is :", diceNumber)

	//dice for LUDO
	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
		// fallthrough para rodar tbm o proximo case caso o dado seja 1
	case 2:
		fmt.Println("you can move 2 spot")
	case 3:
		fmt.Println("You can move 3 spots")
	case 4:
		fmt.Println("You can move 4 spots")
	case 5:
		fmt.Println("You can move 5 spots")
	case 6:
		fmt.Println("You can move 6 spots and roll again!")
	default:
		fmt.Println("What was it!")
	}
}
