package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	hitungDanCetak(a, c)
	hitungDanCetak(b, d)
}

func faktorial(n int) int {
	var hasil int = 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

func permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}

func kombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func hitungDanCetak(n, r int) {
	p := permutasi(n, r)
	k := kombinasi(n, r)
	fmt.Println(p, k)
}