package main

import (
	"fmt"
)

func main() {
	fmt.Println("welcome to slices!")

	var fruitList = []string{"apple", "orange"}
	fmt.Println("type of fruitList is :", len(fruitList))

	fruitList = append(fruitList, "mango", "banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScore := make([]int, 4)

	highScore[0] = 731
	highScore[1] = 935
	highScore[2] = 434
	highScore[3] = 336
	// highScore[4] = 776

	// colcoar mais dados no highscore
	highScore = append(highScore, 555, 666, 667)
	// fmt.Println(highScore)

	// colocar os INT em forma crescente
	// sort.Ints(highScore)
	// fmt.Println(highScore)
	// fmt.Println(sort.IntsAreSorted(highScore))

	//how to remove a value from slices based on index
	// o primeiro index mantem os dados do antes do swift,
	//o segundo index mantem dps do swift, assim apenas o index é removido.
	var courses = []string{"react", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 1
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

	// var fruitList = []string{"apple", "orange", "peach","banana"}
	// var remove int = 1
	// fruitList = append(fruitList[:remove], fruitList[remove+1:]...)
	// fmt.Println(fruitList)

}
