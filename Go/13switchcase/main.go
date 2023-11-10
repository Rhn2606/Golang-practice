package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.NewSource(time.Now().UnixNano())
	diceNum := rand.Intn(6) + 1
	fmt.Println("Value of dice is ", diceNum)
}
