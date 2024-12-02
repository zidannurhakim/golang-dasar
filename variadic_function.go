package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	fmt.Println(sumAll(10, 10, 10))
	fmt.Println(sumAll(10, 10, 10, 20, 50))

	number := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(sumAll(number...))
}
