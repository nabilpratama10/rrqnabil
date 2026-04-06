package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	cetakDeret(n)
}

func cetakDeret(n int) {
	for {
		fmt.Printf("%d ", n)
		
		if n == 1 {
			break
		}

		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
	fmt.Println()
}