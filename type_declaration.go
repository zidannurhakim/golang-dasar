package main

import "fmt"

func main() {
	type NoKTP string

	var ktpZidan NoKTP = "351012"
	var contoh string = "2222222"
	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpZidan)
	fmt.Println(contohKtp)
}
