package main

import "fmt"

func main(){
	var n int
	fmt.Scan(&n)

	if n == 0 {
		fmt.Println("Nol")
	} else {
		p, t := "Ganjil", "Negatif"
		if n%2 == 0 { p = "Genap" }
		if n > 0 { t = "Positif" }
		fmt.Printf("%s, %s\n", p, t)
	}
}
