package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Files in golang")
	content := "this needs to go in a file"

	file, err := os.Create("./mygofile.txt")

	if err != nil {
		panic(err)
	}

	io.WriteString(file, content)

}
