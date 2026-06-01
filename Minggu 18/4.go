package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var pita string
var indeks int
var currentChar byte
var isEOP bool

func start(input string) {
	pita = input
	indeks = 0
	if len(pita) > 0 {
		currentChar = pita[indeks]
		isEOP = (currentChar == '.')
	} else {
		isEOP = true
	}
}

func maju() {
	if indeks < len(pita)-1 {
		indeks++
		currentChar = pita[indeks]
		isEOP = (currentChar == '.')
	} else {
		isEOP = true
	}
}

func eop() bool {
	return isEOP
}

func cc() byte {
	return currentChar
}

func main() {
	fmt.Print("Masukan teks (akhiri dengan titik): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	input = strings.ToUpper(input)

	start(input)

	totalKarakter := 0
	jumlahA := 0
	jumlahLE := 0
	var prevChar byte

	fmt.Print("Karakter yang terbaca: ")
	for !eop() {
		c := cc()
		fmt.Printf("%c", c)

		totalKarakter++

		if c == 'A' {
			jumlahA++
		}

		if prevChar == 'L' && c == 'E' {
			jumlahLE++
		}

		prevChar = c
		maju()
	}
	fmt.Println()

	fmt.Println("Total karakter:", totalKarakter)
	fmt.Println("Jumlah huruf 'A':", jumlahA)

	frekuensi := 0.0
	if totalKarakter > 0 {
		frekuensi = float64(jumlahA) / float64(totalKarakter)
	}
	fmt.Printf("Frekuensi huruf 'A' terhadap total: %.4f\n", frekuensi)
	fmt.Println("Jumlah pasangan kata 'LE':", jumlahLE)
}
