package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var totalTopping int
	fmt.Print("Banyak Topping: ")
	fmt.Scan(&totalTopping)

	toppingDalamPizza := 0
	xc, yc := 0.5, 0.5
	r := 0.5

	for i := 0; i < totalTopping; i++ {
		x := rand.Float64()
		y := rand.Float64()

		jarakKuadrat := (x-xc)*(x-xc) + (y-yc)*(y-yc)

		if jarakKuadrat <= r*r {
			toppingDalamPizza++
		}
	}

	fmt.Printf("Topping pada Pizza: %d\n", toppingDalamPizza)

	pi := 4.0 * float64(toppingDalamPizza) / float64(totalTopping)
	fmt.Printf("PI : %.10f\n", pi)
}
