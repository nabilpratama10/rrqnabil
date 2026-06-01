package main

import (
	"fmt"
	"math/rand"
	"time"
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
	rand.Seed(time.Now().UnixNano())
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

type Pemain struct {
	nama   string
	tangan []Domino
}

func main() {
	var dek Dominoes
	dek.tersisa = 28
	idx := 0
	for i := 0; i <= 6; i++ {
		for j := i; j <= 6; j++ {
			dek.kartu[idx] = Domino{sisi1: i, sisi2: j, nilai: i + j, isBalak: i == j}
			idx++
		}
	}
	kocokKartu(&dek)

	pemain := []Pemain{
		{nama: "Pemain 1"},
		{nama: "Pemain 2"},
		{nama: "Pemain 3"},
		{nama: "Pemain 4"},
	}

	for i := 0; i < 7; i++ {
		for p := 0; p < 4; p++ {
			pemain[p].tangan = append(pemain[p].tangan, ambilKartu(&dek))
		}
	}

	fmt.Println("=== GAME GAPLEH DIMULAI ===")

	ujungKiri := -1
	ujungKanan := -1
	giliran := 0
	passCount := 0
	gameSelesai := false

	for !gameSelesai {
		p := &pemain[giliran]
		kartuDimainkan := false

		for i, k := range p.tangan {
			sisi1 := gambarKartu(k, 1)
			sisi2 := gambarKartu(k, 2)

			if ujungKiri == -1 && ujungKanan == -1 {
				ujungKiri = sisi1
				ujungKanan = sisi2
				fmt.Printf("%s mulai dengan kartu: [%d|%d]\n", p.nama, sisi1, sisi2)
				kartuDimainkan = true
			} else {
				if sisi1 == ujungKiri {
					ujungKiri = sisi2
					fmt.Printf("%s mainin [%d|%d] di kiri. Ujung sekarang: %d dan %d\n", p.nama, sisi1, sisi2, ujungKiri, ujungKanan)
					kartuDimainkan = true
				} else if sisi2 == ujungKiri {
					ujungKiri = sisi1
					fmt.Printf("%s mainin [%d|%d] di kiri. Ujung sekarang: %d dan %d\n", p.nama, sisi1, sisi2, ujungKiri, ujungKanan)
					kartuDimainkan = true
				} else if sisi1 == ujungKanan {
					ujungKanan = sisi2
					fmt.Printf("%s mainin [%d|%d] di kanan. Ujung sekarang: %d dan %d\n", p.nama, sisi1, sisi2, ujungKiri, ujungKanan)
					kartuDimainkan = true
				} else if sisi2 == ujungKanan {
					ujungKanan = sisi1
					fmt.Printf("%s mainin [%d|%d] di kanan. Ujung sekarang: %d dan %d\n", p.nama, sisi1, sisi2, ujungKiri, ujungKanan)
					kartuDimainkan = true
				}
			}

			if kartuDimainkan {
				p.tangan = append(p.tangan[:i], p.tangan[i+1:]...)
				passCount = 0
				break
			}
		}

		if !kartuDimainkan {
			fmt.Printf("%s PASS (Gak ada kartu yang nyambung)\n", p.nama)
			passCount++
		}

		if len(p.tangan) == 0 {
			fmt.Printf("\n🎉 HORE! %s MENANG! Kartunya habis duluan! 🎉\n", p.nama)
			gameSelesai = true
		} else if passCount == 4 {
			fmt.Printf("\n🛑 GAME OVER: GAPLEH! (Mentok, ga ada yang bisa jalan lagi) 🛑\n")
			gameSelesai = true
		}

		giliran = (giliran + 1) % 4
	}
}