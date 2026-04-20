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

type lingkaran struct {
	tengah titik
	radius float64
}

func jarak(p1, p2 titik) float64 {
	return math.Sqrt((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y))
}

func didalam(l lingkaran, p titik) bool {
	return jarak(l.tengah, p) <= l.radius
}

func main() {
	var l1, l2 lingkaran
	var p titik

	fmt.Scan(&l1.tengah.x, &l1.tengah.y, &l1.radius)
	fmt.Scan(&l2.tengah.x, &l2.tengah.y, &l2.radius)
	fmt.Scan(&p.x, &p.y)

	dlm1 := didalam(l1, p)
	dlm2 := didalam(l2, p)

	if dlm1 && dlm2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dlm1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dlm2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}