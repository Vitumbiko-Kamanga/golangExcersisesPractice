package main

import (
	"fmt"

)

func maxMin(m []int) []int{
	var minMax []int
	var min, max int = m[0],m[0]
	for i := 0; i < len(m); i++ {
		if m[i] < min {
			min = m[i]
		} else if m[i] > max {
			max = m[i]
		}else{
			continue
		}
		
	}
	minMax = append(minMax, min)
	minMax = append(minMax, max)
	return minMax
}
func main() {
	grades := []int{90, 81, 77, 73, 70, 84}
	fmt.Println(maxMin(grades))
	
}