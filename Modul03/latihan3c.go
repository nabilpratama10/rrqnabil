package main
import "fmt"

func main() {
	var b int
	var jumlah int

	fmt.Print("Bilangan: ")
	fmt.Scan(&b)

	fmt.Print("Faktor: ")

	for i := 1; i <= b; i++ {
		if b%i == 0 {
			fmt.Print(i, " ")
			jumlah++
		}
	}

	fmt.Println()

	if jumlah == 2 {
		fmt.Println("prima: true")
	} else {
		fmt.Println("prima: false")
	}
}
