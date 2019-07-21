package main

import (
	"fmt"
)

func main() {

	//aritmatika
	x := 9
	y := 7
	z := 5
	//additional
	opTambah := x + y
	fmt.Printf("%v tambah %v adalah %v\n", x, y, opTambah)
	//substracts
	opKurang := opTambah - z
	fmt.Printf("%v kurang %v adalah %v\n", opTambah, z, opKurang)
	//multiplies
	opKali := opKurang * opTambah
	fmt.Printf("%v dikali %v adalah %v\n", opKurang, opTambah, opKali)
	//divides
	opBagi := opKali / 3
	fmt.Printf("%v dibagi 3 adalah %v\n", opKali, opBagi)
	//modulus
	opMod := opKali % 3
	fmt.Printf("%v mod 3 adalah %v\n", opKali, opMod)

	fmt.Println("____________________________________________________")

	//relational operator
	a := 8
	b := 6
	//print relational
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a > b)
	fmt.Println(a < b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)

	fmt.Println("____________________________________________________")

	//logical operator
	adhi := true
	budhi := false
	cindiy := true
	dadanga := false
	//print
	//AND (&&)
	fmt.Println(adhi && budhi)
	fmt.Println(adhi && cindiy)
	fmt.Println(budhi && dadanga)
	//OR (||)
	fmt.Println(adhi || budhi)
	fmt.Println(adhi || cindiy)
	fmt.Println(budhi || dadanga)
	//NOT (!)
	fmt.Println(!adhi)
	fmt.Println(!budhi)
	fmt.Println(!cindiy)
	fmt.Println(!dadanga)
}
