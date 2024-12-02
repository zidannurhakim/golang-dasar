package main

import "fmt"

func main() {
	//Hati hati penggunaan closure karena berdampak ke value counter
	counter := 0

	increment := func() {
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()
	increment()
	fmt.Println(counter)
}
