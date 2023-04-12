package main

import (
	"fmt"
)

func testDataType(variable interface{}) {
	switch variable.(type) {
	case string:
		fmt.Println("It is a string : ", variable.(string))
	case int:
		fmt.Println("It is an integer : ", variable.(int))
	case float64:
		fmt.Println("It is a float : ", variable.(float64))
	default:
		fmt.Println("Data type not present")
	}
}

func main() {
	var variable interface{}
	variable = "Utkarsh"
	testDataType(variable)
	variable = 10
	testDataType(variable)
	variable = 99.9
	testDataType(variable)
	variable = []string{}
	testDataType(variable)
}
