package main

import (
	"fmt"
)

func main() {

	//static var & type data (Regular declarations)
	var x int
	x = 15
	var y float64
	y = 5.5
	//print variable
	fmt.Println(x)
	fmt.Println(y)
	//print var dan type datanya
	fmt.Printf("x type data %T\n", x)
	fmt.Printf("y type data %T\n", y)

	//dynamic var (Short declarations)
	z := "Setyadhi"
	w := 50
	//print variable
	fmt.Println(z)
	fmt.Println(w)
	//print variable dan type data
	fmt.Printf("z type data %T\n", z)
	fmt.Printf("w type data %T\n", w)

	//var aritmatika
	panjang := 5
	lebar := 2
	tinggi := 3
	//print aritmatika
	fmt.Println(panjang * lebar)
	fmt.Println(panjang * lebar * tinggi)
}
