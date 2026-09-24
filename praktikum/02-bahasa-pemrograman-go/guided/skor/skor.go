package main

import "fmt"

func main(){
	var nama string
	var skorMatematika, skorBahasaInggris int

	//Membaca input
	fmt.Scan(&nama)
	fmt.Scan(&skorMatematika)
	fmt.Scan(&skorBahasaInggris)

	// Menghitung total & rata-rata (pembagian bilangan)
	total:= skorMatematika + skorBahasaInggris
	ratarata := total/2

	//Menampilkan output
	fmt.Println(nama)
	fmt.Println(total)
	fmt.Println(ratarata)
}