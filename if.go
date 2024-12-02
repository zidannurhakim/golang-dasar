package main

import "fmt"

func main() {
	name := "Zidan"

	if name == "Zidan" {
		fmt.Println("Halo Zidan")
	} else if name == "Baba" {
		fmt.Println("Sungguh Cute :)")
	} else {
		fmt.Println("Boleh kenalan?")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
