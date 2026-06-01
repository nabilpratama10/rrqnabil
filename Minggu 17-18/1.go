package main

import (
	"fmt"
)

func main() {
	var num, sum float64
	var count int

	for {
		fmt.Scan(&num)
		if num == 9999 {
			break
		}
		sum += num
		count++
	}

	if count > 0 {
		rerata := sum / float64(count)
		fmt.Printf("Rerata: %.2f\n", rerata)
	} else {
		fmt.Println("Belum ada data yang dimasukkan.")
	}
}
