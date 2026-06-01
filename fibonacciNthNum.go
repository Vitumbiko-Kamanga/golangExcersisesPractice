package main

import "fmt"

func fibonacci(n int) int {
	g := 0
	h := 0
	sum := 0
	for i := 0; i <= n; i++ {
		if i == 1 {
			h = 0
			g = 0
		}
		if i == 1 || i == 2 {
			h = 1
			g = 0
		}
		sum = g + h
		g = h
		h = sum
	}
	return sum
}

func main() {
	fmt.Println(fibonacci(10))
}