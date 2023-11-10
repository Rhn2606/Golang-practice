package main

import "fmt"

func main() {

	result := adder(3, 5)

	fmt.Println("Result is ", result)

	proRes := proAddder(2, 5, 4, 6, 3, 6, 4, 8, 7, 8)
	fmt.Println("ProAdder result is", proRes)

}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAddder(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}

	return total
}
