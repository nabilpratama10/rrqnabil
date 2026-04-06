package main

import "fmt"

func cariFaktor(n int, i int) {
	if i <= n {
		if n%i == 0 {
			fmt.Printf("%d ", i)
		}
		cariFaktor(n, i+1)
	}
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan N: ")
	fmt.Scan(&n)

	fmt.Printf("Faktor dari %d adalah: ", n)
	cariFaktor(n, 1)
	fmt.Println()
}