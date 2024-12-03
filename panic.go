package main

import "fmt"

/*
	Func untuk menghentikan program,yang digunakan ketika terjadi panic
	pada saat program kita berjalan namun defer function tetap akan dieksekusi
*/
/*
	Recover adalah function yang dapat digunakan menangkap data panic
	Dengan recover proses panic akan terhenti, sehingga program akan tetap berjalan
*/

func endApp() {
	//Recover yang benar
	message := recover()
	fmt.Println("Terjadi Error", message)
	fmt.Println("Selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Ups Ada Error Panic...")
	}
}

func main() {
	runApp(true)
}
