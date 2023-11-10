package main

import "fmt"

func main() {
	myNum := 23

	var ptr = &myNum

	fmt.Println("Value of pointer is", ptr)
	fmt.Println("Value of pointer is", *ptr)

}
