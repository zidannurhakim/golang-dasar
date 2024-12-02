package main

import "fmt"

func logging() {
	fmt.Println("Selesai Memanggil Function via Defer")
}

func runApps() {
	defer logging()
	fmt.Println("Memulai Menjalankan Apps")
}

func main() {
	runApps()
}
