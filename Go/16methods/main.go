package main

import "fmt"

func main() {

	Rohan := User{"Rohan", "rhn.dev@go", true, 21}
	fmt.Println(Rohan)
	//fmt.Printf("Name is %v and Email is %v", Rohan.Name, Rohan.Email)
	Rohan.GetStatus()
	Rohan.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "testmail@dev.go"
	fmt.Println("Email of this user is: ", u.Email)
}
