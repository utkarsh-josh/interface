package main

import (
	"fmt"
)

func testString(variable interface{}) {
	if val, ok := variable.(string); ok {
		fmt.Println("It is a string : ", val)
		return
	}
	fmt.Println("It is not a string")
}

func testInteger(variable interface{}) {
	if val, ok := variable.(int); ok {
		fmt.Println("It is an integer : ", val)
		return
	}
	fmt.Println("It is not an integer")
}

func testFloat(variable interface{}) {
	if val, ok := variable.(float64); ok {
		fmt.Println("It is a float : ", val)
		return
	}
	fmt.Println("It is not a float")
}

func main() {
	var variable interface{}
	variable = "Utkarsh"
	testString(variable)
	variable = 10
	testInteger(variable)
	variable = 99.9
	testFloat(variable)
	testString(variable)
	testInteger(variable)
}
