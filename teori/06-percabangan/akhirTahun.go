package main

import "fmt"

func main(){
	var belanja int
	var bersedia bool
	fmt.Scan(&belanja, &bersedia)

	kartu := bersedia
	diskon := belanja >= 100000
	cashback := belanja >= 200000 && kartu

	bayar := belanja

	if diskon {
		bayar = bayar - (belanja * 10 / 100)
	}
	if cashback {
		bayar = bayar - 75000
	}
	fmt.Println("kartu?", kartu)
	fmt.Println("diskon?", diskon)
	fmt.Println("cashback?", cashback)
	fmt.Println("Rp", bayar)

}