package main

import "fmt"

func main(){
	var celcius float64

	fmt.Println("Masukkan suhu celcius")
	fmt.Scan(&celcius)

	reamur := celcius * 4/5
	fmt.Println("konversi ke Suhu reamur", reamur)
}