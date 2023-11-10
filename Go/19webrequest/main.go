package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://meta.com"

func main() {
	fmt.Println("Web Request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", response)
	defer response.Body.Close()

	dataBytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	content := string(dataBytes)

	fmt.Println(content)
}
