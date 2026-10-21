package main

import "fmt"

func main(){
	var n, i, s int
	fmt.Println("Masukkan bilangan positif n: ")
	fmt.Scan(&n)
	i = 1
	s = 0

	for i <= n {
		s = s + i
		i++
	}
	fmt.Println("jumlah deret=", s)
}
