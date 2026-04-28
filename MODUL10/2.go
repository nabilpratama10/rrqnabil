package main

import "fmt"

func main() {
	var x, y int
	var berat [1000]float64
	var totalSeluruh, totalWadah float64

	fmt.Scan(&x, &y)

	for i := 0; i < x; i++ {
		fmt.Scan(&berat[i])
		totalSeluruh += berat[i]
	}

	for i := 0; i < x; i++ {
		totalWadah += berat[i]
		
		if (i+1)%y == 0 || i == x-1 {
			fmt.Print(totalWadah, " ")
			totalWadah = 0
		}
	}
	
	fmt.Println()
	fmt.Printf("%.2f\n", totalSeluruh/float64(x))
}