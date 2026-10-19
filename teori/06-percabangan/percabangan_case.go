package main

import "fmt"

func main(){
	lamp_status := "Kunings"

	if lamp_status == "Merah"{
		fmt.Println("Berhenti")
	} else if lamp_status == "Kuning"{
		fmt.Println("Hati-hati / Lebih pelan")
	} else if lamp_status == "Hijau"{
		fmt.Println("Jalan")
	} else{
		fmt.Println("Perintah tidak ada dalam kondisi")
	}
}