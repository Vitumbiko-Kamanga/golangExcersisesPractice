package main
import (
	"fmt"
)
func add(a int, b int) int{
	return a + b
}
func subtract (a int, b int) int {
	return a - b
}

func calculation() {
	var a, b int
	var symbol string
	
	fmt.Scan( &a, &symbol, &b )
	fmt.Print("ANS: ")

	if symbol == "+"{
		fmt.Println(add(a,b))
	}
	if symbol == "-"{
		fmt.Println(subtract(a,b))
	}
	if symbol != "+" && symbol != "-" {
		fmt.Println("Err")
	}
}
func main(){
	fmt.Println("-------------------------")
	fmt.Println("WELCOME TO CLI CALCULATOR")
	fmt.Println("-------------------------")
	calculation()
	fmt.Println("------THANK YOU----------")
}