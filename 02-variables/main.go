package main

import "fmt"

const LoginToken string = "Blablalba" // Public

func main() {
	var username string = "Raphael"
	fmt.Println(username)
	fmt.Printf("Variable type is: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable type is: %T \n", isLoggedIn)

	var smallVal uint32 = 256
	fmt.Println(smallVal)
	fmt.Printf("Variable type is: %T \n", smallVal)

	var smallFloat float64 = 256.4344334344334
	fmt.Println(smallFloat)
	fmt.Printf("Variable type is: %T \n", smallFloat)

	//default values and some aliases
	var anotherVariable int = 1
	fmt.Println(anotherVariable)
	fmt.Printf("Variable type is: %T \n", anotherVariable)

	//implicit type
	var website = "raphael.com"
	fmt.Println(website)

	// no var style
	numberOfUser := 300
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable type is: %T \n", LoginToken)

}
