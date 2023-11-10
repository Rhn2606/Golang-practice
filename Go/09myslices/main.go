package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruits = []string{"apple", "banana", "mango", "pineapple"}
	fmt.Printf("Type of fruits is %T\n", fruits)

	fruits = append(fruits, "AAyien", "Baigan")
	fmt.Println(fruits)

	fruits = append(fruits[1:3])
	fmt.Println(fruits)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 456
	highScores[3] = 867

	highScores = append(highScores, 565, 656, 665)
	sort.Ints(highScores)

	fmt.Println(highScores)

	var courses = []string{"react", "js", "ruby", "go"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
