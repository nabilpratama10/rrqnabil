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
	return d.kartu[d.tersisa]
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

func galiKartu(d *Dominoes, target Domino) {
	for d.tersisa > 0 {
		k := ambilKartu(d)
		if gambarKartu(k, 1) == gambarKartu(target, 1) ||
			gambarKartu(k, 1) == gambarKartu(target, 2) ||
			gambarKartu(k, 2) == gambarKartu(target, 1) ||
			gambarKartu(k, 2) == gambarKartu(target, 2) {
			fmt.Printf("Dapet nih yang nyambung! [%d|%d]\n", gambarKartu(k, 1), gambarKartu(k, 2))
			return
		}
		fmt.Printf("Buang kartu: [%d|%d] (Gak nyambung)\n", gambarKartu(k, 1), gambarKartu(k, 2))
	}
	fmt.Println("Yah, tumpukan abis, ga ada yang nyambung.")
}

func sepasangKartu(k1 Domino, k2 Domino) bool {
	return nilaiKartu(k1)+nilaiKartu(k2) == 12
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

	targetKartu := ambilKartu(&dek)
	fmt.Printf("Kartu patokan di meja: [%d|%d]\n", gambarKartu(targetKartu, 1), gambarKartu(targetKartu, 2))
	fmt.Println("Mulai gali tumpukan...")
	galiKartu(&dek, targetKartu)

	kartuA := Domino{sisi1: 3, sisi2: 4, nilai: 7, isBalak: false}
	kartuB := Domino{sisi1: 2, sisi2: 3, nilai: 5, isBalak: false}
	fmt.Printf("\nCek sepasang kartu A[%d|%d] dan B[%d|%d], totalnya 12 ga? %v\n", gambarKartu(kartuA, 1), gambarKartu(kartuA, 2), gambarKartu(kartuB, 1), gambarKartu(kartuB, 2), sepasangKartu(kartuA, kartuB))
}
