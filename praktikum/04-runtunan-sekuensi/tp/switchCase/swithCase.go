package main

import "fmt"

func main(){
	var hari string
	fmt.Println("Masukkan nama hari:")
	fmt.Scan(&hari)

	switch hari{
	case "Senin", "Rabu" , "Kamis" , "Jumat":
		fmt.Println("Masuk kuliah")
	case "Selasa", "Sabtu" , "Minggu":
		fmt.Println("Libur")
	default:
		fmt.Println("Hari tidak valid")
	}
}