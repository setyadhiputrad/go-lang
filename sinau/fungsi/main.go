package main

import (
	"fmt"
)

func luasKotak(panjang int, lebar int) int {
	return panjang * lebar
}

func userName(nama string) string {
	return fmt.Sprintf("Hello %s", nama)
}

func main() {
	fmt.Println("Sinau fungsi di golang")

	panjang := 7
	lebar := 3
	luas := luasKotak(panjang, lebar)
	fmt.Printf("Luas dari kotak adalah %v\n", luas)

	nama := "setyadhi"
	result := userName(nama)
	fmt.Println(result)

}