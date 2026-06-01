package main
import (
	"fmt"
)

func numberToString(n int) string {
	if n%3 == 0 && n%5 == 0 {
		return "FizzBuzz"
	}
	if n%3 == 0 {
		return "Fizz"
	}
	if n%5 == 0 {
		return "Buzz"
	}
	return fmt.Sprintf("%d", n)
}
func main(){
	for i := 1; i <= 100; i++ {
		fmt.Println(numberToString(i))
	}
}