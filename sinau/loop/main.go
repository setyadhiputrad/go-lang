package main

import (
	"fmt"
)

func main() {

	//standart loop
	for i := 0; i < 5; i++ {
		fmt.Printf("Perulangan ke %d\n", i)
	}

	x := 5
	var y int
	for y < x {
		y++
		fmt.Println(y)
	}

	//break
	a := 1
	for {
		a++
		fmt.Println("hello golang")
		if a == 5 {
			break
		}
	}

	//continue
	b := 1
	for {
		b++
		if b == 3 {
			continue
		}
		fmt.Printf("testing ke %d\n", b)
		if b == 5 {
			break
		}
	}
}
