package main

import "fmt"

func main() {

	loginCount := 15
	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else {
		result = "something else"
	}

	fmt.Println(result)
}
