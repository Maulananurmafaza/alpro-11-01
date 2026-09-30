package main

import "fmt"

func main(){
	var x, y int
	fmt.Println("Jumlah Kue")
	fmt.Scan(&y)
	fmt.Println("Jumlah anggota Keluarga")
	fmt.Scan(&x)

	modulo := y % x
	fmt.Println("Sisa = %d", modulo)
}