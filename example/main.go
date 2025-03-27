package main

import (
	"fmt"
	"github.com/michaelwp/goJsonStruct"
)

func main() {
	jsonConverter := goJsonStruct.NewGoJsonStruct()

	type User struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Email string `json:"email"`
	}

	user := User{Name: "Alice", Age: 30, Email: "alice@example.com"}

	// Convert struct to JSON
	jsonStr := jsonConverter.ToJSON(user)
	fmt.Println(jsonStr) // Output: {"name":"Alice","age":30,"email":"alice@example.com"}

	// Convert struct to pretty JSON
	jsonPretty := jsonConverter.ToJSONIndent(user)
	fmt.Println(jsonPretty)
}
