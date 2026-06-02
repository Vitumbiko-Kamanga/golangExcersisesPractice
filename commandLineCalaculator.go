package main

import (
	"fmt"
	"math"
)
func add(a float64, b float64) float64{
	return a + b
}
func divide (a float64, b float64) float64 {
	if b == 0 {
		fmt.Println("Err")
		return math.Inf(1)
	}
	return a / b
}
func product(a float64, b float64) float64 {
	return a * b
}
func subtract (a float64, b float64) float64 {
	return a - b
}

func calculator() {
	// variables declaration
	var a, b float64
	var symbol string

	// screen formatting 
	fmt.Println("-------------------------")
	fmt.Println("WELCOME TO CLI CALCULATOR")
	fmt.Println("-------------------------")

	// variable entering in the console and reading
	fmt.Scan( &a, &symbol, &b )
	fmt.Print("ANS: ")

	// caling different functions for operations
	if symbol == "+"{
		fmt.Println(add(a,b))
	}
	if symbol == "-"{
		fmt.Println(subtract(a,b))
	}
	if symbol == "/"{
		fmt.Println(divide(a,b))
	}
	if symbol == "*"{
		fmt.Println(product(a,b))
	}
	if symbol != "+" && symbol != "-" && symbol != "*"&& symbol != "/"{
		fmt.Println("Err")
	}
	
	// end of the calculator console or window.
	fmt.Println("------THANK YOU----------")
}

func main(){
	// function calling for a calculator
	calculator()
}