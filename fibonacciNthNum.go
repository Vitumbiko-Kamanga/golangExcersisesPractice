package main

func fibonacci(n int) int{
	g := 0
	h := 1
	sum := 0
	for i := 1; i <= n; i++ {
		if i == 1{
			sum = g
			continue
		}
		if i == 2 {
			sum = h
			continue
		}
		
	}
	return sum
}

func main (){

}