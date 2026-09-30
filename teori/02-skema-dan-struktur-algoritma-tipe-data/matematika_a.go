package main

import "fmt"

func main(){
	var x, y float64
	fmt.Scan(&x, &y)

	f := (1/ (3*x*x +10)) + 10*y + 7
	fmt.Println("Hasilnya adalah ", f)
}