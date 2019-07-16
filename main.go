package main

import (
	"fmt"
)

//membuat konstan publik
const tulisan string = "Ini adalah testing konstan"
const lebar int = 7

func main() {
	//print konstan publik
	fmt.Println(tulisan)

	//membuat konstan lokal
	const panjang int = 16
	//print konstan lokal
	fmt.Println(panjang)
	fmt.Printf("panjang kotak adalah %v\n", panjang)
	fmt.Println(lebar)
	fmt.Printf("lebar kotak adalah %v\n", lebar)

	//example
	z := panjang * lebar
	//print
	fmt.Printf("luas kotak adalah %v\n", z)

}
