package main

import "fmt"

// Muhammad Nabil Raissa P
// 109082500127

type arrInt [256]int

func isiArray(arr *arrInt, n *int) {
	fmt.Scan(n)
	for i := 0; i < *n; i++ {
		fmt.Scan(&arr[i])
	}
}

func reverseArray(arr *arrInt, n int) {
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
	}
}

func isPalindrom(arr arrInt, n int) bool {
	for i := 0; i < n/2; i++ {
		if arr[i] != arr[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var data arrInt
	var n int

	isiArray(&data, &n)

	if isPalindrom(data, n) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}