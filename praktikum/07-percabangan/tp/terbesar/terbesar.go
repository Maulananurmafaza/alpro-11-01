package main

import "fmt"

func main(){
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	terbesar := a

	if b > terbesar {
		terbesar = b
	}
	if c > terbesar {
		terbesar = c
	}
	fmt.Println(terbesar)
}