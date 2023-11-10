package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcoome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza:")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for Rating, ", input)
	fmt.Printf("Type of this rating is %T", input)
}
