package main

import "fmt"

func main() {
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Friday"}
	fmt.Println(days)

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for i, day := range days {
		fmt.Printf("index is %v and value is %v\n", i, day)
	}

	val := 10

	for val < 10 {
		fmt.Println("Value is: ", val)
		val++
	}
}
