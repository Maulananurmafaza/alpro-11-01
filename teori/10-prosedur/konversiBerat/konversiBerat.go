package main

import "fmt"

func konversiBerat(kg float64, gram, ons, pon *int) {
	*gram = int(kg * 1000)
	*ons = *gram / 100
	*pon = *gram / 500
}
func main() {
	var kg float64
	var gram, ons, pon int

	fmt.Scan(&kg)

	konversiBerat(kg, &gram, &ons, &pon)

	fmt.Println(gram, "g")
	fmt.Println(ons, "ons")
	fmt.Println(pon, "pon")
}
