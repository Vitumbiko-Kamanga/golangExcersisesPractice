package main

import (
	"fmt"

)

func main() {
	list := []int{90, 81, 77, 73, 70, 84,100}
	var minMax []int
	var min, max int = list[0], list[0]
	for i := 0; i < len(list); i++ {
		if list[i] < min {
			min = list[i]
		} else if list[i] > max {
			max = list[i]
		}else{
			continue
		}
		
	}
	minMax = append(minMax, min)
	minMax = append(minMax, max)
	fmt.Println(minMax)
}