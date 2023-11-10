package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {

	EncodeJson()
}

func EncodeJson() {
	courses := []course{
		{"ReactJs Bootcamp", 299, "Udemy", "abc23", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 290, "Coursera", "abc43", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "UpGrad", "abc3", []string{"web-dev", "js"}},
	}

	finalJson, err := json.MarshalIndent(courses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func decodeJson() {
	jsonData := []byte(`
	{
			"coursename": "MERN Bootcamp",
            "Price": 290,
            "website": "Coursera",
            "tags": ["full-stack","js"]
	}
	`)
}
