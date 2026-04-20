package main

import "fmt"

// Muhammad Nabil Raissa P
// 109082500127

type Nasabah struct {
	Kode, Nama, Bank string
	Rekening         int
}

func main() {
	var listNasabah [10]Nasabah
	var cariBank string

	for i := 0; i < 10; i++ {
		fmt.Scan(&listNasabah[i].Kode, &listNasabah[i].Nama, &listNasabah[i].Bank, &listNasabah[i].Rekening)
	}

	fmt.Scan(&cariBank)

	for i := 0; i < 10; i++ {
		if listNasabah[i].Bank == cariBank {
			fmt.Printf("Kode: %s , Nasabah: %s , Bank: %s , Rek: %d\n", 
				listNasabah[i].Kode, listNasabah[i].Nama, listNasabah[i].Bank, listNasabah[i].Rekening)
		}
	}
}