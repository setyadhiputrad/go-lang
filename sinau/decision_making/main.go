package main

import (
	"fmt"
)

func ifBolean(state string) string{
	if state {
		fmt.Println("testing decision making")
	}
	return state
}

func perbandingan() {
	panjang := 5
	lebar := 6
	if panjang != lebar {
		fmt.Println("karena nilai panjang lebih kecil dari lebar")
	}
}

func ifLogical() {
	adhi := true
	budi := false
	if adhi || budi {
		fmt.Println("logical OR")
	}
}

//________________________________________________________________________________________________

func ifelse1() {
	diagonal := 59
	if diagonal < 10 {
		fmt.Println("jadi persegi berukuran sedang")
	} else if diagonal >= 59 {
		fmt.Println("jadi persegi berukuran tidak kecil")
	} else {
		fmt.Println("persegi sedang")
	}
}

//_________________________________________________________________________________________________

// expression switch
func expresionSwitch() {
	nilai := 85
	//type_case_1
	switch nilai {
	case 60:
		fmt.Println("Nilai = C")
	case 70:
		fmt.Println("Nilai = B")
	case 85:
		fmt.Println("Nilai = A")
	default:
		fmt.Println("Nilai tidak di ketahui")
	}
	//type_case_2
	switch {
	case nilai == 60:
		fmt.Println("Nilai = C")
	case nilai == 70:
		fmt.Println("Nilai = B")
	case nilai == 85:
		fmt.Println("Nilai = A")
	default:
		fmt.Println("Nilai tidak di ketahui")
	}
}

//type switch
func typeSwitch() {
	var a interface{}
	a = 5
	//
	switch a.(type) {
	case int:
		fmt.Println("a bertipe data integer")
	case string:
		fmt.Println("a bertipe data string")
	case float64:
		fmt.Println("a bertipe data float64")
	default:
		fmt.Println("type data a tidak di ketahui")
	}
}

func main() {

}
