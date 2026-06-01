package main

import (
	"fmt"
	"sort"
)

func main() {
	var data []int
	var num int

	for {
		fmt.Scan(&num)

		if num == -5313 {
			break
		}

		if num == 0 {
			sort.Ints(data)
			n := len(data)

			if n > 0 {
				if n%2 != 0 {
					fmt.Println(data[n/2])
				} else {
					fmt.Println((data[n/2-1] + data[n/2]) / 2)
				}
			}
		} else {
			data = append(data, num)
		}
	}
}
