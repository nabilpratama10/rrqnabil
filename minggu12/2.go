package main

import (
	"fmt"
)

func main() {
	var vote int
	var totalVotes int
	var validVotes int
	var counts [21]int

	for {
		fmt.Scan(&vote)
		if vote == 0 {
			break
		}
		totalVotes++
		if vote >= 1 && vote <= 20 {
			validVotes++
			counts[vote]++
		}
	}

	fmt.Printf("Suara masuk: %d\n", totalVotes)
	fmt.Printf("Suara sah: %d\n", validVotes)

	max1 := 0
	max2 := 0
	ketua := 0
	wakil := 0

	for i := 1; i <= 20; i++ {
		if counts[i] > max1 {
			max2 = max1
			wakil = ketua
			max1 = counts[i]
			ketua = i
		} else if counts[i] > max2 {
			max2 = counts[i]
			wakil = i
		}
	}

	fmt.Printf("Ketua RT: %d\n", ketua)
	fmt.Printf("Wakil ketua: %d\n", wakil)
}
