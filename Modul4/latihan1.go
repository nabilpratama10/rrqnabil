package main

import "fmt"

func main() {
	var bilangan int
	var pesan string

	fmt.Scan(&bilangan, &pesan)
	cetakPesan(pesan, bilangan)
}

func cetakPesan(M string, flag int) {
	var jenis string

	switch flag {
	case 0:
		jenis = "error"
	case 1:
		jenis = "warning"
	case 2:
		jenis = "informasi"
	default:
		jenis = "unknown"
	}

	fmt.Println(M, jenis)
}