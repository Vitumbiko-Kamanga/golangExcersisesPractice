package main

import (
    "fmt"
    //"math/rand"
    //"math"
)
/*
// func sum(n []float64) float64{
// total := 0.0
// for _,v := range n{
//     total +=v
// }
// return total
// }

func avg(n []float64) float64{
    av := 0.0
    total := sum(n)
    av = total / 6.0
    return av
}
*/
func square( n int) int{
    return n*n
    
}


func oldNums(n int) (int){
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
    fmt.Println(oldNums(5))


/*
    //  fname := "vitumbiko"
    //  greeting(fname)

    //  //calling the sum function to demonstrate
    //  grades := []float64 {80, 74, 72, 76, 83, 79} // this is the slice 
    //  fmt.Printf("the sum of the grades is %v \n",sum(grades))
    //  fmt.Printf("the average of your grades is %0.2f \n", avg(grades))
     
    //  //generation of a random integer between 0-10
    //  fmt.Println("My random number is ", rand.Intn(10))
    //  // suing the math library
    //  fmt.Println("here is the use of a math Library; ", math.Sqrt(16))
    //  fmt.Printf("here is the use of a math Librar:  %0.3f \n", math.Pi)
*/
}
