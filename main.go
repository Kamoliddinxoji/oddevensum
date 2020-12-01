package main

import "fmt"

func main() {
	var num int
	oddSum := 0
	evenSum := 0
	var i = 1

	fmt.Scan(&num)

	for i <= num {
		if i%2 == 0 {
			evenSum += i
		} else {
			oddSum += i
		}
		i++
	}

	fmt.Printf("Sum of Odd number = %d  and even number %d \n", oddSum, evenSum)

}
