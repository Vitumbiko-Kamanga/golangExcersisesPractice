package main

import (
    "fmt"
)
// the function for perfoming the number squaring
func square( n int) int{
    return n*n
    
}

// function for generating the sum of Old Perfect Square numbers
func oldPerfectSquareNums(n int) (int){
    m := []int {}
    sum := 0
    for i := 1; i <= n; i++ {
        num := square(i)
        if num%2 != 0{
                sum +=num
                m = append(m, num)
        } 
    }
    return sum
}

func main() {
    // function calling for finding the sum of Old Perfect Square numbers in a given range (parameter)
    fmt.Println(oldPerfectSquareNums(5))

}
