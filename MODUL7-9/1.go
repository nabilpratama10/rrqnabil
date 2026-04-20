package main

import (
	"fmt"
	"math"
)

// Muhammad Nabil Raissa P
// 109082500127

type titik struct {
	x, y float64
}

func jarak(P1, P2 titik) float64 {
	return math.Sqrt((P1.x-P2.x)*(P1.x-P2.x) + (P1.y-P2.y)*(P1.y-P2.y))
}

func main() {
	var p1, p2 titik

	fmt.Scan(&p1.x, &p1.y, &p2.x, &p2.y)

	fmt.Println(jarak(p1, p2))
}