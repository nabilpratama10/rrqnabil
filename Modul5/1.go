package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	baris(n)
}

func baris(bilangan int) {
	if bilangan <= 0 {
		return
	}

	fmt.Println(bilangan)

	if bilangan > 1 {
		baris(bilangan - 1)
	}
}