package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("N suku pertama: ")
	fmt.Scan(&n)

	piPrev := 0.0
	piCurr := 0.0

	for i := 1; i <= n; i++ {
		if i > 1 {
			diff := math.Abs(piCurr - piPrev)
			if diff <= 0.00001 {
				fmt.Printf("Hasil PI: %.10f\n", piPrev)
				fmt.Printf("Hasil PI: %.10f\n", piCurr)
				fmt.Printf("Pada i ke: %d\n", i)
				return
			}
		}

		term := 1.0 / float64(2*i-1)
		if i%2 == 0 {
			term = -term
		}

		piPrev = piCurr
		piCurr += 4 * term
	}

	fmt.Printf("Hasil PI: %.7f\n", piCurr)
}
