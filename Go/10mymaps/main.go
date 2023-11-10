package main

import "fmt"

func main() {
	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of languages: ", languages)
	fmt.Println("JS shorts for: ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("New List of languages: ", languages)

	//loops

	for key, value := range languages {
		fmt.Printf("For key %v,  Value is %v\n", key, value)
	}
}
