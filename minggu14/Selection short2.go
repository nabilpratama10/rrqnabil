package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)

		arr := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&arr[j])
		}

		for j := 0; j < m-1; j++ {
			minIdx := j
			for k := j + 1; k < m; k++ {
				if arr[k] < arr[minIdx] {
					minIdx = k
				}
			}
			arr[j], arr[minIdx] = arr[minIdx], arr[j]
		}

		first := true
		for j := 0; j < m; j++ {
			if arr[j]%2 != 0 {
				if !first {
					fmt.Print(" ")
				}
				fmt.Print(arr[j])
				first = false
			}
		}

		for j := m - 1; j >= 0; j-- {
			if arr[j]%2 == 0 {
				if !first {
					fmt.Print(" ")
				}
				fmt.Print(arr[j])
				first = false
			}
		}
		fmt.Println()
	}
}
