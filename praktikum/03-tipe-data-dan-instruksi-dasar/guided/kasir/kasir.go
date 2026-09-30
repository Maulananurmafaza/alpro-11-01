package main

import "fmt"

func main(){
	var x int 
	fmt.Println("Masukan nominal")
	fmt.Scan(&x)

	var sepuluhRibuan int = x / 10000
	var sisa int = x % 10000

	var limaRibuan int = sisa / 5000
	sisa = sisa % 5000

	var seRibuan int = sisa / 1000

	fmt.Println(sepuluhRibuan, limaRibuan, seRibuan)
}