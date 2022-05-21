package main

import "fmt"

func main() {
	fmt.Println("welcome to my structs!:)~")
	// no inheritance in golang, no super or parent

	type User struct {
		Name   string
		Email  string
		Status bool
		Age    int
	}

	raphael := User{"Raphael", "raphael@go.dev", true, 29}
	fmt.Println(raphael)
	fmt.Printf("raphael details are: %+v\n", raphael)
	fmt.Printf("Name is %v and email is %v.", raphael.Name, raphael.Email)

}
