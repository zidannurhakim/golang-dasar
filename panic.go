package main

import "fmt"

// Func untuk menghentikan program,
// yang digunakan ketika terjadi panic pada saat program kita berjalan namun defer function tetap akan dieksekusi

func endApp() {
	fmt.Println("Selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Ups Ada Error..")
	}
}

func main() {
	runApp(false)
}
