package main

import "fmt"

func main() {
	var nama, pemenangNama string
	var soal, skor, maxSoal, minSkor int

	maxSoal = -1
	minSkor = 999999

	for {
		fmt.Scan(&nama)
		if nama == "Selesai" {
			break
		}

		var currentSoal, currentSkor int
		hitungSkor(&currentSoal, &currentSkor)

		if currentSoal > maxSoal {
			maxSoal = currentSoal
			minSkor = currentSkor
			pemenangNama = nama
		} else if currentSoal == maxSoal {
			if currentSkor < minSkor {
				minSkor = currentSkor
				pemenangNama = nama
			}
		}
	}

	fmt.Printf("%s %d %d\n", pemenangNama, maxSoal, minSkor)
}

func hitungSkor(soal *int, skor *int) {
	var waktu int
	*soal = 0
	*skor = 0

	for i := 0; i < 8; i++ {
		fmt.Scan(&waktu)
		if waktu <= 300 {
			*soal++
			*skor += waktu
		}
	}
}