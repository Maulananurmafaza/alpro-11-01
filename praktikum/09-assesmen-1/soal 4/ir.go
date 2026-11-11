package main

import "fmt"


func prima(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
func main (){
	var L, R, jumlah int
	
	fmt.Scan(&L, &R)

	for i := L; i <= R; i++ {
		a := i 
		total := 0
		sebelumnya := 0
		valid := true

		for a > 0 {
			digit := a % 10

			if digit < sebelumnya {
				valid = false
			}
			sebelumnya = digit
			total += digit
			a /= 10
		}
		if valid && prima (total) {
			jumlah++
		}
	}
	fmt.Println(jumlah)
}
	
