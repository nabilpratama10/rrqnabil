package main

import (
	"fmt"
)

func main() {
	var x string
	var n int

	fmt.Scan(&x)
	fmt.Scan(&n)

	data := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}

	var found bool
	var positions []int
	var count int

	for i := 0; i < n; i++ {
		if data[i] == x {
			found = true
			positions = append(positions, i+1)
			count++
		}
	}

	if found {
		fmt.Println("a. Apakah string x ada? Ya.")
	} else {
		fmt.Println("a. Apakah string x ada? Tidak.")
	}

	if found {
		fmt.Print("b. Ditemukan pada posisi ke: ")
		for i, pos := range positions {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Print(pos)
		}
		fmt.Println()
	} else {
		fmt.Println("b. Ditemukan pada posisi ke: -")
	}

	fmt.Printf("c. Ada berapakah string x? %d buah.\n", count)

	if count >= 2 {
		fmt.Println("d. Adakah sedikitnya dua string x? Ya.")
	} else {
		fmt.Println("d. Adakah sedikitnya dua string x? Tidak.")
	}
}
