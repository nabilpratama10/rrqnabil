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
	for i := 1; i <= 20; i++ {
		if counts[i] > 0 {
			fmt.Printf("%d : %d\n", i, counts[i])
		}
	}
}
