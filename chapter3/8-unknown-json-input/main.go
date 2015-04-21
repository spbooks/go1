package main

import (
	"encoding/json"
	"fmt"
)

func FooJSON(input string) {
	data := map[string]interface{}{}
	err := json.Unmarshal([]byte(input), &data)
	if err != nil {
		panic(err)
	}
	foo, _ := data["foo"]
	switch foo.(type) {
	case float64:
		fmt.Printf("Float %f\n", foo)
	case string:
		fmt.Printf("String %s\n", foo)
	default:
		fmt.Printf("Something else\n")
	}
}
func main() {
	FooJSON(`{
		"foo": 123
	}`)
	FooJSON(`{
		"foo": "bar"
	}`)
	FooJSON(`{
		"foo": []
	}`)
}
