package main

import (
	"fmt"
	"strings"

	"github.com/ashkenazi1/requester"
)

func main() {
	//Get Request
	r := requester.New()
	resp, err := r.Get("https://www.example.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Response Status:", string(resp))

	// Post Request
	t := requester.New()
	body := strings.NewReader(`{"name": "John", "age": 30}`)
	res, err := t.Post("https://jsonplaceholder.typicode.com/posts", body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Response:", string(res))

	// PostJson Request
	c := requester.New()
	body2 := map[string]interface{}{
		"name": "John",
		"age":  30,
	}
	respo, err := c.PostJSON("https://jsonplaceholder.typicode.com/posts", body2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Response:", string(respo))
}
