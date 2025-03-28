package goJsonStruct

import (
	"encoding/json"
	"fmt"
)

type GoJsonStruct interface {
	ToJSON(v interface{}) string
	ToStruct(jsonStr string, v interface{}) error
	ToJSONPretty(v interface{}) string
}

type ImplGoJsonStruct struct{}

func NewGoJsonStruct() GoJsonStruct {
	return ImplGoJsonStruct{}
}

// ToJSON converts any struct to a JSON string.
func (j ImplGoJsonStruct) ToJSON(v interface{}) string {
	jsonData, err := json.Marshal(v)
	if err != nil {
		return "{}" // Return empty JSON if there's an error
	}
	return string(jsonData)
}

// ToJSONPretty converts any struct to a prettified JSON string.
func (j ImplGoJsonStruct) ToJSONPretty(v interface{}) string {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "{}"
	}
	return string(jsonData)
}

// ToStruct converts a JSON string back to a struct.
func (j ImplGoJsonStruct) ToStruct(jsonStr string, v interface{}) error {
	if err := json.Unmarshal([]byte(jsonStr), v); err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}
	return nil
}
