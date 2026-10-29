package main

import "fmt"

func main(){
	var jenis string
	var jam, tarif int
	fmt.Scan(&jenis, &jam)

	if jenis == "mobil" {
		tarif = 5000 + (jam-1)*3000
		if jam > 24 {
			tarif += 50000
		}
	} else if jenis == "motor" {
		tarif = 3000 + (jam-1)*2000
		if jam > 24 {
			tarif += 30000
		}
	}
	fmt.Println(tarif)
}