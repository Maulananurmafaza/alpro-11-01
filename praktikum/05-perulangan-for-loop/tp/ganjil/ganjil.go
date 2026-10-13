package main

import "fmt"

func main(){
	var n int
	jumlah := 0

	fmt.Println("Masukkan bilangan positif: ")
	fmt.Scan(&n)

	fmt.Println("Bilangan ganjil: ")


	for i := 1; i <= n; i+= 2 {
		fmt.Println(i, " ")
		jumlah +=i
	}

	fmt.Println()
	fmt.Println("jumlah =", jumlah)
}