package main

import "fmt"

func main() {
	fmt.Println("welcome to my structs!:)~")
	// no inheritance in golang, no super or parent
	raphael := User{"Raphael", "raphael@go.dev", true, 29}
	fmt.Println(raphael)
	fmt.Printf("raphael details are: %+v\n", raphael)
	fmt.Printf("Name is %v and email is %v.\n", raphael.Name, raphael.Email)
	raphael.GetStatus()
	raphael.NewMail()
	fmt.Printf("Name is %v and email is %v.\n", raphael.Name, raphael.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
