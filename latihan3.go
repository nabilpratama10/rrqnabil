package main

import "fmt"

func main() {
	var r int
	const pi = 3.1415926535

	fmt.Print("Jari-jari: ")
	fmt.Scan(&r)

	volume := (4.0 / 3.0) * pi * float64(r*r*r)
	luas := 4 * pi * float64(r*r)

	fmt.Printf("Volume bola = %.4f\n", volume)
	fmt.Printf("Luas bola = %.4f\n", luas)
}