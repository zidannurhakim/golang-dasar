package main

import "fmt"

type Address struct {
	City, Province, Country string
}

type Address1 struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Banyuwangi", "Jawa Timur", "Indonesia"}
	address2 := address1

	address2.City = "Malang"
	fmt.Println(address1) //Tidak Berubah
	fmt.Println(address2) //Baru berubah

	addressKe1 := Address{"Banyuwangi", "Jawa Timur", "Indonesia"}
	addressKe2 := &addressKe1 //Pointer

	addressKe2.City = "Jember"
	fmt.Println(addressKe1) //Berubah karena fitur pointer &
	fmt.Println(addressKe2) //Berubah

}
