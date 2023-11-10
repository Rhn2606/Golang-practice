package main

import "fmt"

func main() {

	type User struct {
		Name   string
		Email  string
		Status bool
		Age    int
	}

	Rohan := User{"Rohan", "rhn.dev@go", true, 21}
	fmt.Println(Rohan)
	fmt.Printf("Name is %v and Email is %v", Rohan.Name, Rohan.Email)
}
