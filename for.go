package main

import "fmt"

func main() {
	//MODEL 1
	//counter := 1
	//
	//for counter <= 10 {
	//	fmt.Println("Perulangan ke", counter)
	//	counter++
	//}
	//fmt.Println("Selesai")

	//MODEL 2
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}
	fmt.Println("Selesai")

	//FOR RANGE
	//MANUAL
	names := []string{"Alex", "John", "Jane"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, name := range names {
		fmt.Println("Index", index, "=", name)
	}

	for _, name := range names {
		fmt.Println(name)
	}
}
