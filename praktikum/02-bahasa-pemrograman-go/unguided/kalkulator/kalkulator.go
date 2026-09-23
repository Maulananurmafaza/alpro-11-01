package main

import "fmt"

func main() {
	var a,b int

	//membaca input
	fmt.Print("masukan nilai a")
	fmt.Scan(&a)
	fmt.Print("masukan nilai b")
	fmt.Scan(&b)

	//menampilkan output
	fmt.Print("Hasil Penjumlahan: ")
	fmt.Println(a + b)
	fmt.Print("Hasil Pengurangan: ")
	fmt.Println(a - b)
	fmt.Print("Hasil Perkalian: ")
	fmt.Println(a * b)
	fmt.Print("Hasil Pembagian: ")
	fmt.Println(a / b)
	fmt.Print("Hasil Sisa Hasil Bagi: ")
	fmt.Println(a % b)

}
