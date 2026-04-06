package main

import "fmt"

func urutanBolakBalik(n int) {
	if n == 1 {
		fmt.Printf("%d ", n)
		return
	}
	fmt.Printf("%d ", n)
	urutanBolakBalik(n - 1)
	fmt.Printf("%d ", n)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n)
	urutanBolakBalik(n)
	fmt.Println()
}