package main

import "fmt"

//“If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
//Find the sum of all the multiples of 3 or 5 below 1000.”

func main() {

	sum := 0

	for temp := 1; temp <= 1000; temp++ {
		if temp%3 == 0 || temp%5 == 0 {
			sum += temp
		}
	}

	fmt.Println("total sum of all multiple of 3 or 5 below 1000 is :", sum)
}
