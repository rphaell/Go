package main

import "fmt"

func main() {
	fmt.Println("Learning arrays in go")

	var fruitList [4]string

	fruitList[0] = "apple"
	fruitList[1] = "orange"
	fruitList[3] = "banana"

	//imprimir a array
	fmt.Println("fruit list is: ", fruitList)
	//imprimir o length da array
	fmt.Println("fruit list is: ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushrooms"}
	fmt.Println("vegList is: ", len(vegList))
}
