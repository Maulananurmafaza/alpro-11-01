package main

import "fmt"

func main(){
	var hari int
	var tahun, bulan, minggu, sisaHari int

	fmt.Println("masukan jumlah Hari")
	fmt.Scan(&hari)

	tahun = hari / 360
	hari = hari % 360
	bulan = hari / 30 
	hari = hari % 30
	minggu = hari / 7
	sisaHari = hari % 7

	fmt.Println("tahun", tahun)
	fmt.Println("bulan", bulan)
	fmt.Println("minggu", minggu)
	fmt.Println("sisaHari", sisaHari)


}