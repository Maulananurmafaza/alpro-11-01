package main

import "fmt"

func main(){
	var angka, sebelumnya int
	var total, jumlah int
	var tertinggi, terendah int

	fmt.Scan(&angka)

	tertinggi = angka 
	terendah = angka
	total = angka 
	jumlah = 1
	sebelumnya = angka

	for {
		fmt.Scan(&angka)
		if angka == sebelumnya {
			break
		}
		if angka > tertinggi {
			tertinggi = angka
		}
		if angka < terendah {
			terendah = angka
		}
		total += angka
		jumlah++

		sebelumnya = angka
	}
	fmt.Printf("%d %d %.2f\n", tertinggi, terendah, float64(total) / float64(jumlah))

}