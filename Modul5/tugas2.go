package main

import (
	"fmt"
	"strings"
)

func printStars(n int) {
	if n > 0 {
		printStars(n - 1)
		fmt.Println(strings.Repeat("*", n))
	}
}

func main() {
	var n int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n)

	printStars(n)
}