package main

import (
	"fmt"
	"math/rand"
)

type Domino struct {
	sisi1   int
	sisi2   int
	nilai   int
	isBalak bool
}

type Dominoes struct {
	kartu   [28]Domino
	tersisa int
}

func kocokKartu(d *Dominoes) {
	for i := 0; i < d.tersisa; i++ {
		j := rand.Intn(d.tersisa)
		d.kartu[i], d.kartu[j] = d.kartu[j], d.kartu[i]
	}
}

func ambilKartu(d *Dominoes) Domino {
	if d.tersisa == 0 {
		return Domino{}
	}
	d.tersisa--
	kartuDiambil := d.kartu[d.tersisa]
	return kartuDiambil
}

func gambarKartu(d Domino, suit int) int {
	if suit == 1 {
		return d.sisi1
	}
	return d.sisi2
}

func nilaiKartu(d Domino) int {
	return d.nilai
}

func main() {
	var dek Dominoes
	dek.tersisa = 28

	idx := 0
	for i := 0; i <= 6; i++ {
		for j := i; j <= 6; j++ {
			dek.kartu[idx] = Domino{
				sisi1:   i,
				sisi2:   j,
				nilai:   i + j,
				isBalak: i == j,
			}
			idx++
		}
	}

	kocokKartu(&dek)

	kartuKu := ambilKartu(&dek)

	fmt.Println("Kartu yang diambil:")
	fmt.Println("Sisi 1:", gambarKartu(kartuKu, 1))
	fmt.Println("Sisi 2:", gambarKartu(kartuKu, 2))
	fmt.Println("Nilai Total:", nilaiKartu(kartuKu))
	fmt.Println("Apakah Balak?", kartuKu.isBalak)
	fmt.Println("Sisa kartu di dek:", dek.tersisa)
}
